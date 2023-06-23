resource "aws_route53_resolver_query_log_config_association" "pike" {
  resolver_query_log_config_id = aws_route53_resolver_query_log_config.example.id
  resource_id                  = aws_vpc.example.id
}