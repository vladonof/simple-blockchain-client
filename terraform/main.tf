module "vpc" {
  source = "./modules/vpc"

  cidr_block = "10.0.0.0/16"
}

module "ecr" {
  source = "./modules/ecr"

  service_name = var.service_name
}

module "cluster" {
  source = "./modules/cluster"

  cluster_name = var.cluster_name
}

module "lb" {
  source = "./modules/lb"

  service_name = var.service_name
  region       = var.region
  vpc_id       = module.vpc.vpc_id
  lb_subnets   = module.vpc.public_subnets

  lb_listener_port            = var.lb_listener_port
  lb_target_group_port        = var.lb_target_group_port
  lb_target_group_healthcheck = var.lb_target_group_healthcheck
}

module "service" {
  source = "./modules/service"

  service_name               = var.service_name
  region                     = var.region
  ecs_desired_count          = var.ecs_desired_count
  ecs_subnets                = module.vpc.public_subnets
  ecs_container_name         = var.service_name
  ecs_container_port         = var.ecs_container_port
  ecs_task_definition_cpu    = var.ecs_task_definition_cpu
  ecs_task_definition_memory = var.ecs_task_definition_memory
  ecs_cluster                = module.cluster.cluster_id
  lb_target_group_arn        = module.lb.target_group_arn
  ecr_image_tag              = "latest"
  ecr_image_repository_url   = module.ecr.repository_url
}
