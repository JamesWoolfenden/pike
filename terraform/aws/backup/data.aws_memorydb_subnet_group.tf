data "aws_memorydb_subnet_group" "pike" {
  name = "pike"
}

output "aws_memorydb_subnet_group" {
  value = data.aws_memorydb_subnet_group.pike
}
