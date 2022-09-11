data "aws_cognito_user_pool_client" "client" {
  client_id    = "f9ubuthmv61ppp6f0pksqocu9"
  user_pool_id = "eu-west-2_vC5Ib0JRX"
}

output "thisclient" {
  sensitive = true
  value     = data.aws_cognito_user_pool_client.client
}
