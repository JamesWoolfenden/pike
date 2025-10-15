data "aws_odb_db_servers" "pike" {
  cloud_exadata_infrastructure_id = "pike"
}

output "aws_odb_db_servers" {
  value = data.aws_odb_db_servers.pike
}
