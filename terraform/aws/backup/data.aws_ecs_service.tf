data "aws_ecs_service" "pike" {
  cluster_arn  = ""
  service_name = "pike"
}
