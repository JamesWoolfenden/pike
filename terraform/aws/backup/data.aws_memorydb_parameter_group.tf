data "aws_memorydb_parameter_group" "pike" {
  name = "pike"
}

output "aws_memorydb_parameter_group" {
  value = data.aws_memorydb_parameter_group.pike
}
