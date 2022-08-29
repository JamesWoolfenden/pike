resource "aws_codeartifact_domain" "pike" {
  domain = "pike"
  tags = {
    pike = "permissions"
  }
}
