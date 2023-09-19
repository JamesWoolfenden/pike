data "aws_cognito_user_pools" "selected" {
  name = "pike"
}

output "user_pools" {
  value = data.aws_cognito_user_pools.selected
}
