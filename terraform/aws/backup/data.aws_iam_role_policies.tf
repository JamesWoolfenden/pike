data "aws_iam_role_policies" "pike" {
  role_name = "pike"
}

output "aws_iam_role_policies" {
  value = data.aws_iam_role_policies.pike
}
