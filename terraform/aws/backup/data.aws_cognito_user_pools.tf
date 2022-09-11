data "aws_cognito_user_pools" "selected" {
  name = "pike"
}

output "pools" {
  value = data.aws_cognito_user_pools.selected
}
