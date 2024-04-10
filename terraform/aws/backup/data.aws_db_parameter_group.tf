data "aws_db_parameter_group" "pike" {
  name = "default.postgres15"
}

output "aws_db_parameter_group" {
  value = data.aws_db_parameter_group.pike
}
