resource "aws_cognito_resource_server" "pike" {
  identifier = "https://example.com"
  name       = "example"

  scope {
    scope_name        = "sample-scope"
    scope_description = "a Sample Scope Description"
  }

  user_pool_id = "eu-west-2_vC5Ib0JRX"
}
