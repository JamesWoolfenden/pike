data "aws_odb_cloud_autonomous_vm_cluster" "pike" {
}

output "aws_odb_cloud_autonomous_vm_cluster" {
  value = data.aws_odb_cloud_autonomous_vm_cluster.pike
}
