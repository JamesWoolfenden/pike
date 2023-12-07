resource "aws_route53_resolver_firewall_config" "pike" {
  resource_id        = data.aws_vpc.pike.id
  firewall_fail_open = "ENABLED"
}
