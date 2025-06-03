resource "aws_notifications_event_rule" "pike" {
  event_pattern = jsonencode({
    detail = {
      state = {
        value = ["ALARM"]
      }
    }
  })
  event_type                     = "CloudWatch Alarm State Change"
  notification_configuration_arn = aws_notifications_notification_configuration.pike.arn
  regions                        = ["us-east-1", "us-west-2"]
  source                         = "aws.cloudwatch"
}
