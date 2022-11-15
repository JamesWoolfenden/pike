resource "aws_cloudwatch_composite_alarm" "pike" {
  alarm_description = "This is a composite alarm!"
  alarm_name        = "example-composite-alarm-pike"

  alarm_actions = ["arn:aws:sns:eu-west-2:680235478471:pike"]
  ok_actions    = ["arn:aws:sns:eu-west-2:680235478471:pike"]

  alarm_rule = "ALARM(root_usage_alarm) OR ALARM(unauthorized_api_calls_alarm)"
  tags = {
    pike = "permissions"
  }
}
