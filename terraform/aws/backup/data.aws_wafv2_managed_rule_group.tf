data "aws_wafv2_managed_rule_group" "pike" {
  name  = "some-rule-group"
  scope = "REGIONAL"
}

output "aws_wafv2_managed_rule_group" {
  value = data.aws_wafv2_managed_rule_group.pike
}
