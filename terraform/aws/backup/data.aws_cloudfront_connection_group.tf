data "aws_cloudfront_connection_group" "pike" {
  id = "pike-12345"
}

output "aws_cloudfront_connection_group" {
  value = data.aws_cloudfront_connection_group.pike
}
