
resource "aws_guardduty_member_detector_feature" "runtime_monitoring" {
  detector_id = aws_guardduty_detector.example.id
  account_id  = "680235478471"
  name        = "RUNTIME_MONITORING"
  status      = "ENABLED"

  additional_configuration {
    name   = "EKS_ADDON_MANAGEMENT"
    status = "ENABLED"
  }

  additional_configuration {
    name   = "ECS_FARGATE_AGENT_MANAGEMENT"
    status = "ENABLED"
  }
}
