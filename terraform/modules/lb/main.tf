resource "aws_lb" "this" {
  load_balancer_type = "application"
  name               = "${var.service_name}-lb"

  subnets = var.lb_subnets
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.this.arn
  port              = var.lb_listener_port
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.arn
  }
}

resource "aws_lb_target_group" "this" {
  name        = "${var.service_name}-http"
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_id
  port        = var.lb_target_group_port
  health_check {
    path                = var.lb_target_group_healthcheck.path
    protocol            = var.lb_target_group_healthcheck.protocol
    matcher             = var.lb_target_group_healthcheck.matcher
    interval            = var.lb_target_group_healthcheck.interval
    timeout             = var.lb_target_group_healthcheck.timeout
    healthy_threshold   = var.lb_target_group_healthcheck.healthy_threshold
    unhealthy_threshold = var.lb_target_group_healthcheck.unhealthy_threshold
  }
}
