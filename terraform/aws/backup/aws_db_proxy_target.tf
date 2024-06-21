resource "aws_db_proxy_target" "pike" {
  db_proxy_name          = aws_db_proxy.pike.name
  target_group_name      = aws_db_proxy_default_target_group.pike.name
  db_instance_identifier = aws_db_instance.pike.identifier
}
