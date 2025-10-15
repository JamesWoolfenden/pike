data "aws_odb_db_nodes" "pike" {
}

output "aws_odb_db_nodes" {
  value = data.aws_odb_db_nodes.pike
}
