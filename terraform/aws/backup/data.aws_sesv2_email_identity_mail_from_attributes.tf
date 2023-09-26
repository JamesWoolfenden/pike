data "aws_sesv2_email_identity_mail_from_attributes" "pike" {
  //email_identity = data.aws_sesv2_email_identity.pike.email_identity
  email_identity = "freddy"
}

output "identity" {
  value = data.aws_sesv2_email_identity_mail_from_attributes.pike
}
