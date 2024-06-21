resource "aws_db_proxy_endpoint" "pike" {
  db_proxy_name          = aws_db_proxy.pike.name
  db_proxy_endpoint_name = "example"
  vpc_subnet_ids         = data.aws_subnets.public.ids
  target_role            = "READ_ONLY"
}
