data "aws_odb_cloud_autonomous_vm_clusters" "pike" {
}

output "aws_odb_cloud_autonomous_vm_clusters" {
  value = data.aws_odb_cloud_autonomous_vm_clusters.pike
}
