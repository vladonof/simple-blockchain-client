variable "cluster_name" {
  type    = string
  default = "APIs"
}

variable "service_name" {
  type    = string
  default = "simple-blockchain-client"
}

variable "namespace" {
  type    = string
  default = "trustwallet"
}

variable "region" {
  type    = string
  default = "eu-west-x"
}

variable "environment" {
  type    = string
  default = "staging"
}

variable "lb_listener_port" {
  type    = number
  default = 8080
}

variable "lb_target_group_port" {
  type        = number
  description = "Target group port ( port that the conatiners listens on )"
  default     = 8080
}

variable "lb_target_group_healthcheck" {
  description = "Target group healthcheck config"
  type = object({
    path                = optional(string)
    protocol            = optional(string)
    matcher             = optional(string)
    interval            = number
    timeout             = number
    healthy_threshold   = number
    unhealthy_threshold = number
  })
  default = {
    path                = "/healthcheck"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }
}

variable "ecs_desired_count" {
  type        = number
  description = "Number of instances of the task definition to place and keep running."
  default     = 1
}

variable "ecs_container_port" {
  type        = number
  description = "Port on the container to associate with the load balancer"
  default     = 8080
}

variable "ecs_task_definition_cpu" {
  type        = number
  description = "CPU allocated to the task definition container"
  default     = 256
}

variable "ecs_task_definition_memory" {
  type        = number
  description = "Memory allocated to the task definition container"
  default     = 512
}
