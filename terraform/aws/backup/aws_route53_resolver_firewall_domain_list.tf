resource "aws_route53_resolver_firewall_domain_list" "pike" {
  name = "pike"
  tags = {
    pike = "permissions"
  }
}
