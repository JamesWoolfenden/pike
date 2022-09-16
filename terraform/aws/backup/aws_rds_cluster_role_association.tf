resource "aws_rds_cluster_role_association" "pike" {
  db_cluster_identifier = "aurora-cluster-pike"
  feature_name          = "S3_INTEGRATION"
  role_arn              = "arn:aws:iam::680235478471:role/rds-monitoring-role"
}
