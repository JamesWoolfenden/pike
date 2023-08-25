resource "aws_route53_resolver_rule_association" "pike" {
  resolver_rule_id = aws_route53_resolver_rule.sys.id
  vpc_id           = aws_vpc.foo.id
}
