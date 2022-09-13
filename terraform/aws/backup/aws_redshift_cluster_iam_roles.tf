resource "aws_redshift_cluster_iam_roles" "pike" {
  cluster_identifier = "redshift-cluster-1"
  iam_role_arns      = ["arn:aws:iam::680235478471:role/aws-service-role/redshift.amazonaws.com/AWSServiceRoleForRedshift"]
}
