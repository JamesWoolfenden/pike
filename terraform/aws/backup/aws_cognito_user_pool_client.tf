resource "aws_cognito_user_pool_client" "client" {
  name            = "client"
  generate_secret = true
  user_pool_id    = "eu-west-2_vC5Ib0JRX"
}
