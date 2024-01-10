resource "aws_servicecatalog_service_action" "pike" {
  description = "Motor generator unit"
  name        = "MGU"

  definition {
    name    = "AWS-RestartEC2Instance"
    version = "1"
  }
}
