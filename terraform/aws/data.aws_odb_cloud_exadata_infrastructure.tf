data "aws_odb_cloud_exadata_infrastructure" "pike" {
}

output "aws_odb_cloud_exadata_infrastructure" {
  value = data.aws_odb_cloud_exadata_infrastructure.pike
}
