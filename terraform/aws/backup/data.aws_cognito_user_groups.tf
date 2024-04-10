data "aws_cognito_user_groups" "pike" {
  user_pool_id = "us-west-2_aaaaaaaaa"
}


output "aws_cognito_user_groups" {
  value = data.aws_cognito_user_groups.pike
}
