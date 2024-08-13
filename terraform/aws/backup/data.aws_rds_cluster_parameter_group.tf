data "aws_rds_cluster_parameter_group" "pike" {
  name = "pike"
}

output "aws_rds_cluster_parameter_group" {
  value = data.aws_rds_cluster_parameter_group.pike
}
