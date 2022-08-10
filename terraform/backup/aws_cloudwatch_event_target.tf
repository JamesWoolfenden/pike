resource "aws_cloudwatch_event_rule" "stop_instances" {
  name                = "StopInstance"
  description         = "Stop instances nightly"
  schedule_expression = "cron(0 0 * * ? *)"
}

resource "aws_cloudwatch_event_target" "stop_instances" {
  target_id = "StopInstance"
  arn       = "arn:aws:ssm:eu-west-2::document/AWS-RunShellScript"
  input     = "{\"commands\":[\"halt\"]}"
  rule      = aws_cloudwatch_event_rule.stop_instances.name
  role_arn  = "arn:aws:iam::680235478471:role/pike_ssm_lifecycle"

  run_command_targets {
    key    = "tag:Terminate"
    values = ["midnight"]
  }
}
