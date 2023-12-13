resource "aws_route53_resolver_dnssec_config" "pike" {
  resource_id = data.aws_vpc.pike.id
}
