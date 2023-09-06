data "aws_ecs_task_execution" "pike" {
  cluster         = "pike"
  task_definition = "arn:aws:ecs:eu-west-2:680235478471:task-definition/fargate-task-jp-stg:1"
}
