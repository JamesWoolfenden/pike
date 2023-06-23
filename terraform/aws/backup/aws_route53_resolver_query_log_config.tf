resource "aws_route53_resolver_query_log_config" "pike" {
  name            = "example"
  destination_arn = aws_s3_bucket.example.arn

  tags = {
    Environment = "Prod"
  }
}