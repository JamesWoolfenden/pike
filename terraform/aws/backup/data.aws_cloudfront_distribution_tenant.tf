data "aws_cloudfront_distribution_tenant" "pike" {
  name = "pike"
}

output "aws_cloudfront_distribution_tenant" {
  value = data.aws_cloudfront_distribution_tenant.pike
}
