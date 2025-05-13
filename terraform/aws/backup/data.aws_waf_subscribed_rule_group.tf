data "aws_waf_subscribed_rule_group" "pike" {
  name = "F5 Bot Detection Signatures For AWS WAF"
}

output "aws_waf_subscribed_rule_group" {
  value = data.aws_waf_subscribed_rule_group.pike
}
