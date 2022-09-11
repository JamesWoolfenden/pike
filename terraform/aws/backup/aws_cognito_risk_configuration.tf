resource "aws_cognito_risk_configuration" "pike" {
  user_pool_id = "eu-west-2_vC5Ib0JRX"

  risk_exception_configuration {
    blocked_ip_range_list = ["10.10.10.10/32"]
  }
}
