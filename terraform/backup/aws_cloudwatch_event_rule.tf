resource "aws_cloudwatch_event_rule" "console" {
  name        = "capture-aws-sign-in"
  description = "Capture each AWS Console Sign In"

  event_pattern = <<EOF
{
  "detail-type": [
    "AWS Console Sign In via CloudTrail"
  ]
}
EOF

  #  schedule_expression="cron(0 20 * * ? *)"
  #  event_bus_name=
  role_arn = "arn:aws:iam::680235478471:role/pike-test-role"
  tags = {
    test = "permissions"
  }
}
