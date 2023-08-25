resource "aws_ses_identity_policy" "pike" {
  identity = aws_ses_domain_identity.example.arn
  name     = "example"
  policy   = data.aws_iam_policy_document.example.json
}
