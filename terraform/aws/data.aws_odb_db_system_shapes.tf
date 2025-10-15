data "aws_odb_db_system_shapes" "pike" {
}

output "aws_odb_db_system_shapes" {
  value = data.aws_odb_db_system_shapes.pike
}
