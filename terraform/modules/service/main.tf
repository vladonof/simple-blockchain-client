resource "aws_ecs_task_definition" "this" {
  family                   = "${var.service_name}-task-def"
  execution_role_arn       = aws_iam_role.this.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory

  container_definitions = templatefile("${path.module}/taskdef.json", {
    SERVICE_NAME  = var.service_name
    ECR_ADDRESS   = var.ecr_image_repository_url
    ECR_IMAGE_TAG = var.ecr_image_tag
    AWS_REGION    = var.region
  })
}

resource "aws_ecs_service" "this" {
  task_definition                   = aws_ecs_task_definition.this.arn
  name                              = var.service_name
  cluster                           = var.ecs_cluster
  desired_count                     = var.ecs_desired_count
  launch_type                       = "FARGATE"
  enable_ecs_managed_tags           = true
  health_check_grace_period_seconds = 60
  wait_for_steady_state             = false


  network_configuration {
    assign_public_ip = true
    subnets          = var.ecs_subnets
  }

  load_balancer {
    target_group_arn = var.lb_target_group_arn
    container_name   = var.ecs_container_name
    container_port   = var.ecs_container_port
  }
}


resource "aws_iam_role" "this" {
  name = "${var.service_name}-task-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "this" {
  role       = aws_iam_role.this.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_cloudwatch_log_group" "this" {
  name = "${var.service_name}-group"
}

resource "aws_cloudwatch_log_stream" "this" {
  name           = "${var.service_name}-stream"
  log_group_name = aws_cloudwatch_log_group.this.name
}
