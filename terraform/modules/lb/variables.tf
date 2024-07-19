variable "service_name" {
  type     = string
  nullable = false
}

variable "region" {
  type     = string
  nullable = false
}


variable "vpc_id" {
  type     = string
  nullable = false
}

variable "lb_subnets" {
  type     = list(string)
  nullable = false
}

variable "lb_listener_port" {
  type     = number
  nullable = false
}

variable "lb_target_group_port" {
  type        = number
  description = "Target group port ( port that the conatiners listens on )"
  nullable    = false
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
  nullable = false
}
