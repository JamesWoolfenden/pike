resource "aws_route53_resolver_config" "pike" {
  autodefined_reverse_flag = "DISABLE"
  resource_id              = data.aws_vpc.pike.id
}
