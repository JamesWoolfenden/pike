resource "aws_redshift_usage_limit" "pike" {
  cluster_identifier = "redshift-cluster-1"
  feature_type       = "concurrency-scaling"
  limit_type         = "time"
  amount             = 70

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
