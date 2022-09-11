resource "aws_cognito_user" "pike" {
  user_pool_id             = "eu-west-2_vC5Ib0JRX"
  username                 = "example"
  desired_delivery_mediums = ["EMAIL", "SMS"]
}

output "user" {
  sensitive = true
  value     = aws_cognito_user.pike
}
