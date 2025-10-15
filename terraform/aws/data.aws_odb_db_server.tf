data "aws_odb_db_server" "pike" {
}

output "aws_odb_db_server" {
  value = data.aws_odb_db_server.pike
}
