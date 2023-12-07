resource "aws_route53_resolver_firewall_rule_group_association" "pike" {
  vpc_id                 = data.aws_vpc.pike.id
  firewall_rule_group_id = aws_route53_resolver_firewall_rule_group.pike.id
  name                   = "pike"
  priority               = 101
}

data "aws_vpc" "pike" {
  id = "vpc-06074a092930bc809"
}
