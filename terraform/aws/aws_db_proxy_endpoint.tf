#resource "aws_db_proxy_endpoint" "pike" {
#  db_proxy_name          = aws_db_proxy.pike.name
#
#  db_proxy_endpoint_name = "example"
#  vpc_subnet_ids         = aws_subnet.test.*.id
#  target_role            = "READ_ONLY"
#tags = {
#  pike="permissions"
#  delete="me"
#}
#}
