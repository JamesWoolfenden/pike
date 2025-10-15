data "aws_odb_db_node" "pike" {
  cloud_vm_cluster_id = "pike"
  id                  = "pike"
}

output "aws_odb_db_node" {
  value = data.aws_odb_db_node.pike
}
