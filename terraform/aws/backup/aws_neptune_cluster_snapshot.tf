resource "aws_neptune_cluster_snapshot" "pike" {
  db_cluster_identifier          = aws_neptune_cluster.pike.id
  db_cluster_snapshot_identifier = "resourcetestsnapshot1234"
}
