resource "aws_inspector2_filter" "pike" {
  name   = "pike"
  action = "NONE"
  filter_criteria {
    aws_account_id {
      comparison = "EQUALS"
      value      = "111222333444"
    }
  }

  tags = {
    pike = "permission"
  }
}
