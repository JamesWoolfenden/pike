resource "aws_appconfig_environment" "pike" {
  name           = "example-environment-tf"
  description    = "Example AppConfig Environment"
  application_id = aws_appconfig_application.example.id

  monitor {
    alarm_arn      = aws_cloudwatch_metric_alarm.example.arn
    alarm_role_arn = aws_iam_role.example.arn
  }

  tags = {
    Type = "AppConfig Environment"
  }
}
