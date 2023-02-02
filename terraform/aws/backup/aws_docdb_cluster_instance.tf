resource "aws_docdb_cluster_instance" "pike" {
  count              = 2
  identifier         = "docdb-cluster-demo-${count.index}"
  cluster_identifier = aws_docdb_cluster.pike.id
  instance_class     = "db.r5.large"
}
