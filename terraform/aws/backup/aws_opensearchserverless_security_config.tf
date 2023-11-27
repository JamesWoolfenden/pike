resource "aws_opensearchserverless_security_config" "pike" {
  provider = aws.central
  name     = "pike"
  type     = "saml"
  saml_options {
    metadata = file("${path.module}/idp-metadata.xml")
  }

}
