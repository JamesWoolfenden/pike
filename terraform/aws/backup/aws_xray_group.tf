resource "aws_xray_group" "pike" {
  group_name        = "pike"
  filter_expression = "responsetime > 5"

  insights_configuration {
    insights_enabled      = true
    notifications_enabled = true
  }
  tags = {
    pike = "permissions"
  }
}
