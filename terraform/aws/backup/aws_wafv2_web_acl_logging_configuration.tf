resource "aws_wafv2_web_acl_logging_configuration" "pike" {
  resource_arn            = aws_wafv2_web_acl.example.arn
  log_destination_configs = ["arn:aws:logs:eu-west-2:680235478471:log-group:pike/"]
}

resource "aws_cloudwatch_log_group" "pike" {
  name = "pike"
}
