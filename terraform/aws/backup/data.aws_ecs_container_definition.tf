data "aws_ecs_container_definition" "pike" {
  container_name  = "pike"
  task_definition = "mongodb"
}
