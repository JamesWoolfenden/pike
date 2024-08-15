resource "aws_cognito_identity_pool_provider_principal_tag" "pike" {
  identity_pool_id       = aws_cognito_identity_pool.example.id
  identity_provider_name = aws_cognito_user_pool.example.endpoint
  use_defaults           = false
  principal_tags = {
    test = "value"
  }
}
