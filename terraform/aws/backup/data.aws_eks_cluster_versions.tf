data "aws_eks_cluster_versions" "pike" {
}

output "aws_eks_cluster_versions" {
  value = data.aws_eks_cluster_versions.pike
}
