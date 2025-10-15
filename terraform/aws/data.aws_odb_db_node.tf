data "aws_odb_db_node" "pike" {
}

output "aws_odb_db_node" {
  value = data.aws_odb_db_node.pike
}
