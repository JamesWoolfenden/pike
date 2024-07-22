data "aws_cloudfront_origin_access_control" "pike" {
  id = "E2T5VTFBZJ3BJB"
}

output "aws_cloudfront_origin_access_control" {
  value = data.aws_cloudfront_origin_access_control.pike
}
