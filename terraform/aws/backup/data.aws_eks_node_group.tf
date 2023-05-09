data "aws_eks_node_group" "pike" {
  cluster_name    = "pike"
  node_group_name = "pike"
}
