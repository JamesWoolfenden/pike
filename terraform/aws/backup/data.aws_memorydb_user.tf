data "aws_memorydb_user" "pike" {
  user_name = "pike"
}

output "aws_memorydb_user" {
  value = data.aws_memorydb_user.pike
}
