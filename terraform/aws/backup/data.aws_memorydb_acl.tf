data "aws_memorydb_acl" "pike" {
  name = "pike"
}

output "aws_memorydb_acl" {
  value = data.aws_memorydb_acl.pike
}
