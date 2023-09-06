data "aws_ecrpublic_authorization_token" "pike" {}

output "token" {
  value     = data.aws_ecrpublic_authorization_token.pike
  sensitive = true
}
