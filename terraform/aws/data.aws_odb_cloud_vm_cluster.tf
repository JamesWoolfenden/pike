data "aws_odb_cloud_vm_cluster" "pike" {
}

output "aws_odb_cloud_vm_cluster" {
  value = data.aws_odb_cloud_vm_cluster.pike
}
