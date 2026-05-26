data "aws_iam_outbound_web_identity_federation" "pike" {
}

output "aws_iam_outbound_web_identity_federation" {
  value = data.aws_iam_outbound_web_identity_federation.pike
}
