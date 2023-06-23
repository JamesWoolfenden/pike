resource "aws_ses_domain_dkim" "pike" {
  domain = aws_ses_domain_identity.example.domain
}