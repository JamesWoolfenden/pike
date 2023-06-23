resource "aws_ses_domain_identity_verification" "pike" {
  domain = aws_ses_domain_identity.example.id
}