resource "aws_db_cluster_snapshot" "pike" {
  db_cluster_identifier          = "aurora-cluster-pike"
  db_cluster_snapshot_identifier = "pike1234"
  tags = {
    pike = "permission"
    //delete="me"
  }
}
