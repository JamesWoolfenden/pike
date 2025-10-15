data "aws_odb_cloud_exadata_infrastructure" "pike" {
  id = "pike"
}

output "aws_odb_cloud_exadata_infrastructure" {
  value = data.aws_odb_cloud_exadata_infrastructure.pike
}
