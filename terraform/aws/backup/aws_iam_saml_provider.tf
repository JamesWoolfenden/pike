resource "aws_iam_saml_provider" "default" {
  name                   = "myprovider"
  saml_metadata_document = file("saml-metadata.xml")
}
