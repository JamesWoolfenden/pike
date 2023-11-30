
resource "aws_cloudfront_distribution" "staging" {
  staging = true
  default_cache_behavior {
    allowed_methods        = ["POST", "HEAD", "PATCH", "DELETE", "PUT", "GET", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "privateforcloudfront.s3.eu-west-2.amazonaws.com"
    viewer_protocol_policy = "allow-all"
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6"
    compress               = true
  }
  origin {
    connection_attempts = 3
    connection_timeout  = 10
    domain_name         = "privateforcloudfront.s3.eu-west-2.amazonaws.com"
    origin_id           = "privateforcloudfront.s3.eu-west-2.amazonaws.com"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
  viewer_certificate {
    cloudfront_default_certificate = true
  }
  http_version    = "http2"
  enabled         = true
  is_ipv6_enabled = true
  tags = {
    pike = "permission"
  }
}

data "aws_s3_bucket" "selected" {
  bucket = "privateforcloudfront"
}




resource "aws_cloudfront_continuous_deployment_policy" "example" {
  enabled = true

  staging_distribution_dns_names {
    items    = [aws_cloudfront_distribution.staging.domain_name]
    quantity = 1
  }

  traffic_config {
    type = "SingleWeight"
    single_weight_config {
      weight = "0.01"
    }
  }
}
