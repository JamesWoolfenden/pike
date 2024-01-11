resource "aws_codestarconnections_connection" "example" {
  name          = "example-connection"
  provider_type = "Bitbucket"
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
