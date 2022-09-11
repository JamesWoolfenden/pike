resource "aws_cognito_user_group" "main" {
  name         = "pike"
  user_pool_id = "eu-west-2_vC5Ib0JRX"
  description  = "Managed by Terraform file"
  precedence   = 42
  role_arn     = "arn:aws:iam::680235478471:role/duffcognitotrust"
}
