data "aws_cognito_user_pool" "pike" {
  user_pool_id = "us-west-2_aaaaaaaaa"
}

output "aws_cognito_user_pool" {
  value = data.aws_cognito_user_pool.pike
}
