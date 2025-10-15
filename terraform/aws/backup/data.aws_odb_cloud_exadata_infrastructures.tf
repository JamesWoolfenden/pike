data "aws_odb_cloud_exadata_infrastructures" "pike" {
}

output "aws_odb_cloud_exadata_infrastructures" {
  value = data.aws_odb_cloud_exadata_infrastructures.pike
}
