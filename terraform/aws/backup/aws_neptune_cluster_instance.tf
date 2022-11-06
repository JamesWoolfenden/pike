resource "aws_neptune_cluster_instance" "pike" {
  identifier                   = "pike"
  instance_class               = "db.t3.medium"
  neptune_parameter_group_name = "example"
  engine                       = "neptune"
  apply_immediately            = true
  auto_minor_version_upgrade   = true
  neptune_subnet_group_name    = "main"
  cluster_identifier           = aws_neptune_cluster.pike.cluster_identifier
  tags = {
    pike = "permissions"
  }
}
