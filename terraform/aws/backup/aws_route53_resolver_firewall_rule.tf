resource "aws_route53_resolver_firewall_rule" "pike" {
  action                  = "ALLOW"
  firewall_domain_list_id = aws_route53_resolver_firewall_domain_list.pike.id
  firewall_rule_group_id  = aws_route53_resolver_firewall_rule_group.pike.id
  name                    = "pike"
  priority                = 100
}
