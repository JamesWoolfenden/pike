resource "aws_cloudfront_public_key" "pike" {
  comment     = "pike update"
  encoded_key = data.tls_public_key.pike.public_key_pem
  name        = "pike"
}
