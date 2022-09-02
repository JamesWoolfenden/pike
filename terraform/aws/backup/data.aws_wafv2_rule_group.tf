data "aws_wafv2_rule_group" "pike" {
  name  = "global"
  scope = "CLOUDFRONT"
}

output "aws_wafv2_rule_group" {
  value = data.aws_wafv2_rule_group.pike
}
