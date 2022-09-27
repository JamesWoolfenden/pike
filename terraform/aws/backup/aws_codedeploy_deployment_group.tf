resource "aws_codedeploy_deployment_group" "pike" {
  app_name               = "pike"
  deployment_group_name  = "pike"
  service_role_arn       = "arn:aws:iam::680235478471:role/codedeployservice"
  deployment_config_name = aws_codedeploy_deployment_config.pike.id

  trigger_configuration {
    trigger_events     = ["DeploymentFailure"]
    trigger_name       = "example-trigger"
    trigger_target_arn = "arn:aws:sns:eu-west-2:680235478471:pike"
  }
}
