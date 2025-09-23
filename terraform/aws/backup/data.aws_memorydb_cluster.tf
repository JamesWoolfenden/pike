data "aws_memorydb_cluster" "pike" {
  name = "pike"
}

output "aws_memorydb_cluster" {
  value = data.aws_memorydb_cluster.pike
}
