data "aws_odb_db_server" "pike" {
  id                              = "pike"
  cloud_exadata_infrastructure_id = "pike"
}

output "aws_odb_db_server" {
  value = data.aws_odb_db_server.pike
}
