data "aws_cloudwatch_contributor_managed_insight_rules" "pike" {
  resource_arn = "arn:aws:ec2:us-west-2:680235478471:resource-name/resourceid"
}

output "aws_cloudwatch_contributor_managed_insight_rules" {
  value = data.aws_cloudwatch_contributor_managed_insight_rules.pike
}
