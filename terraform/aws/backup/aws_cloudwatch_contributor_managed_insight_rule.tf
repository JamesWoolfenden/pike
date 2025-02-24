resource "aws_cloudwatch_contributor_managed_insight_rule" "pike" {
  resource_arn  = "arn:aws:logs:eu-west-2:680235478471:log-group:/aws/connect/pike2"
  template_name = "VpcEndpointService-BytesByEndpointId-v2"
  # rule_state    = "DISABLED"
}
