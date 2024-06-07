resource "aws_ses_active_receipt_rule_set" "pike" {
  rule_set_name = aws_ses_receipt_rule_set.example.id
}

resource "aws_ses_receipt_rule_set" "example" {
  rule_set_name = "example"
}
