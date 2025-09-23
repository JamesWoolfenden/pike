data "aws_memorydb_snapshot" "pike" {
  name = "pike"
}

output "aws_memorydb_snapshot" {
  value = data.aws_memorydb_snapshot.pike
}
