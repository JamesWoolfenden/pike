resource "aws_route53_resolver_firewall_rule_group" "pike" {
  name = "pike"
  tags = {
    pike = "permissions"
  }
}
