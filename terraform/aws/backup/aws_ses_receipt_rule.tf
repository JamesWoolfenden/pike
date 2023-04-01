resource "aws_ses_receipt_rule" "example" {
  name          = "example"
  rule_set_name = "example"
  enabled       = true
  scan_enabled  = true
  tls_policy    = "Require"

  recipients = ["example@example.com"]
}
