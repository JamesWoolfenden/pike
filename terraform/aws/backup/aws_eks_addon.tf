resource "aws_eks_addon" "pike" {
  cluster_name = aws_eks_cluster.pike.name
  addon_name   = "vpc-cni"
}
