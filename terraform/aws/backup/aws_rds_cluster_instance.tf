resource "aws_rds_cluster_instance" "pike" {
  identifier         = "aurora-cluster-demo"
  cluster_identifier = "aurora-cluster-pike"
  instance_class     = "db.t2.small"
  #  engine             = aws_rds_cluster.default.engine
  #  engine_version     = aws_rds_cluster.default.engine_version
  monitoring_role_arn = "arn:aws:iam::680235478471:role/rds-monitoring-role"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
  engine = ""
}
