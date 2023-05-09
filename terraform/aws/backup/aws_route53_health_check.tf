resource "aws_route53_health_check" "example" {
  fqdn              = "example.com"
  port              = 443
  type              = "HTTPS"
  resource_path     = "/health"
  failure_threshold = 2
  request_interval  = 30
  measure_latency   = true

  tags = {
    pike = "permissions"
  }
}
