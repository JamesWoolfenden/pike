resource "aws_cloudfront_origin_access_control" "pike" {
  name                              = "pike"
  description                       = "Example Policy"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}
