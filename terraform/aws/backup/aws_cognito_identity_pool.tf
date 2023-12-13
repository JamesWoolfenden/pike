resource "aws_cognito_identity_pool" "pike" {
  identity_pool_name = "identity pool"

  allow_unauthenticated_identities = false
  allow_classic_flow               = false
  tags = {
    pike    = "permissions"
    another = "one"
  }
}
