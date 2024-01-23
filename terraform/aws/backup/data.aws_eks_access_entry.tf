data "aws_eks_access_entry" "pike" {
  principal_arn = "arn:aws:cloudformation:us-east-1:123456789012:changeSet/MyProductionChangeSet/abc9dbf0-43c2-11e3-a6e8-50fa526be49c"
  cluster_name  = "pike"
}
