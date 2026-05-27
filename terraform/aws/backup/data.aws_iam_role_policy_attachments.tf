data "aws_iam_role_policy_attachments" "pike" {
  role_name = "pike"
}

output "aws_iam_role_policy_attachments" {
  value = data.aws_iam_role_policy_attachments.pike
}
