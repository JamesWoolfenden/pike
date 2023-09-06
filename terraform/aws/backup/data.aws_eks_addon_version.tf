data "aws_eks_addon_version" "pike" {
  kubernetes_version = "1.17"
  addon_name         = "pike"
}
