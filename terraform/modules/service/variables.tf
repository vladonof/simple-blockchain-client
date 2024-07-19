variable "service_name" {
  type     = string
  nullable = false
}

variable "region" {
  type     = string
  nullable = false
}

variable "ecs_desired_count" {
  type        = number
  description = "Number of instances of the task definition to place and keep running."
  nullable    = false
}

variable "ecs_subnets" {
  type        = list(string)
  description = "Subnets where the aws_ecs_service resource is deployed"
  nullable    = false
}

variable "ecs_container_name" {
  type        = string
  description = "Name of the container to associate with the load balancer ( as it appaears in the task def )"
  nullable    = false
}

variable "ecs_container_port" {
  type        = number
  description = "Port on the container to associate with the load balancer"
  nullable    = false
}

variable "ecs_task_definition_cpu" {
  type        = number
  description = "CPU allocated to the task definition container"
  nullable    = false
}

variable "ecs_task_definition_memory" {
  type        = number
  description = "Memory allocated to the task definition container"
  nullable    = false
}

variable "ecs_cluster" {
  type        = string
  description = "ECS cluster id"
  nullable    = false
}

variable "lb_target_group_arn" {
  type        = string
  description = "Load balancer target group arn"
  nullable    = false
}

variable "ecr_image_tag" {
  type        = string
  description = "The tag of the image to use"
  nullable    = false
}

variable "ecr_image_repository_url" {
  type        = string
  description = "The URL of ECR repository"
  nullable    = false
}
