resource "aws_ses_domain_mail_from" "pike" {
  domain           = aws_ses_domain_identity.example.domain
  mail_from_domain = "bounce.${aws_ses_domain_identity.example.domain}"
}
