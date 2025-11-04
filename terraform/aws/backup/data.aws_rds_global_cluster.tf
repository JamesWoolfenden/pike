data "aws_rds_global_cluster" "pike" {
  identifier = "pike"
}

output "aws_rds_global_cluster" {
  value = data.aws_rds_global_cluster.pike
}
