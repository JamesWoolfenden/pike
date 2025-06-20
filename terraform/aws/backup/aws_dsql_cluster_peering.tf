resource "aws_dsql_cluster_peering" "pike" {
  clusters       = [aws_dsql_cluster.pike.arn]
  identifier     = "test-peer"
  witness_region = "eu-west-2"
}
#
# resource "aws_dsql_cluster" "two" {}
