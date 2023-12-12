resource "aws_cloudfront_function" "pike" {
  name    = "test"
  runtime = "cloudfront-js-1.0"
  comment = "my new function b"
  publish = true
  code    = file("${path.module}/function.js")
}
