resource "aws_route53_resolver_rule" "pike" {
  domain_name = "subdomain.freebeer.site"
  rule_type   = "SYSTEM"
  tags = {
    pike = "permissions"
  }
}
