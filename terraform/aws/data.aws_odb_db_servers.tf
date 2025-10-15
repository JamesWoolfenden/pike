data "aws_odb_db_servers" "pike" {
}

output "aws_odb_db_servers" {
  value = data.aws_odb_db_servers.pike
}
