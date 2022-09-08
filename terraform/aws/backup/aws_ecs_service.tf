resource "aws_ecs_service" "mongo" {
  name            = "mongodb"
  cluster         = aws_ecs_cluster.pike.id
  task_definition = aws_ecs_task_definition.service.arn
  desired_count   = 3
  // iam_role        = "arn:aws:iam::680235478471:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS"


  ordered_placement_strategy {
    type  = "binpack"
    field = "cpu"
  }

  #  load_balancer {
  #    target_group_arn = aws_lb_target_group.foo.arn
  #    container_name   = "mongo"
  #    container_port   = 8080
  #  }

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [eu-west-2a, eu-west-2b]"
  }

  network_configuration {
    subnets = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7", "subnet-05a6a6de2f4989d22"]
  }

  tags = {
    pike = "permissions"
  }
}
