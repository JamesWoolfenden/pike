data "aws_cognito_user_pool_clients" "main" {
  user_pool_id = "eu-west-2_vC5Ib0JRX"
}

output "clients" {
  value = data.aws_cognito_user_pool_clients.main
}
