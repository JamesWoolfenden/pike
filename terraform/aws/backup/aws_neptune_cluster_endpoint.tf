resource "aws_neptune_cluster_endpoint" "pike" {
  cluster_identifier          = aws_neptune_cluster.pike.cluster_identifier
  cluster_endpoint_identifier = "pike"
  endpoint_type               = "READER"
  tags = {
    pike = "permissions"
  }
}
