data "aws_redshift_cluster_credentials" "pike" {
  cluster_identifier = "pike"
  db_user            = "pike"
}
