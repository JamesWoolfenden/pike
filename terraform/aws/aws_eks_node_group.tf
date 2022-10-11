resource "aws_eks_node_group" "pike" {
  cluster_name  = aws_eks_cluster.pike.name
  node_role_arn = "arn:aws:iam::680235478471:role/eks-node-group20221011093427965800000001"
  subnet_ids    = ["subnet-09ff91b5b0adb1fd4", "subnet-05e87623a2a5c41f0", "subnet-043bb893867355740"]

  node_group_name_prefix = aws_eks_cluster.pike.name
  scaling_config {
    desired_size = 1
    max_size     = 2
    min_size     = 1
  }

  update_config {
    max_unavailable = 1
  }
  tags = {
    pike = "permissions"
  }
}
