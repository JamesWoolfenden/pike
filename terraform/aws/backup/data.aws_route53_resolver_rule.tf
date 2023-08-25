data "aws_route53_resolver_rule" "pike" {
  domain_name = "subdomain.example.com"
  rule_type   = "SYSTEM"
}
