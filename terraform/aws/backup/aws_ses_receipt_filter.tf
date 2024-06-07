resource "aws_ses_receipt_filter" "pike" {
  name   = "block-spammer"
  cidr   = "10.10.10.11"
  policy = "Block"
}
