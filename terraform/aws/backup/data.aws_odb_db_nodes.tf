data "aws_odb_db_nodes" "pike" {
  cloud_vm_cluster_id = "pike"
}

output "aws_odb_db_nodes" {
  value = data.aws_odb_db_nodes.pike
}
