resource "aws_inspector_assessment_template" "pike" {
  name       = "pike"
  target_arn = aws_inspector_assessment_target.pike.arn
  duration   = 4600

  rules_package_arns = data.aws_inspector_rules_packages.pike.arns

  event_subscription {
    event     = "ASSESSMENT_RUN_COMPLETED"
    topic_arn = "arn:aws:sns:eu-west-2:680235478471:assessement"
  }

  tags = {
    pike = "permissions"
    // delete="me"
  }
}
