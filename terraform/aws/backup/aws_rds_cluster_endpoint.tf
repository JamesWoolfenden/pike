resource "aws_rds_cluster_endpoint" "pike" {
  cluster_identifier          = "aurora-cluster-pike"
  cluster_endpoint_identifier = "static"
  custom_endpoint_type        = "READER"

  static_members = [
    aws_rds_cluster_instance.pike.id,
    // aws_rds_cluster_instance.pike2.id,
  ]

  #tags = {
  #  pike="permission"
  #  delete="me"
  #}
}
