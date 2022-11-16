resource "aws_cloudfront_key_group" "pike" {
  comment = "example key group pike"
  items   = [aws_cloudfront_public_key.pike.id]
  name    = "example-key-group"
}
