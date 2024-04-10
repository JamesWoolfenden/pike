data "aws_cognito_user_group" "pike" {
  user_pool_id = "us-west-2_aaaaaaaaa"
  name         = "example"
}

output "aws_cognito_user_group" {
  value = data.aws_cognito_user_group.pike
}
