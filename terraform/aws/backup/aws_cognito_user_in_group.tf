resource "aws_cognito_user_in_group" "pike" {
  group_name   = "pike"
  user_pool_id = "eu-west-2_vC5Ib0JRX"
  username     = "example"
}
