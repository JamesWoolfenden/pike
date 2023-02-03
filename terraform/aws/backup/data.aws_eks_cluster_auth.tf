data "aws_eks_cluster_auth" "pike" {
  name = "pike"
}

output "data" {
  value     = data.aws_eks_cluster_auth.pike
  sensitive = true
}
