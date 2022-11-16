resource "aws_cloudfront_monitoring_subscription" "pike" {
  distribution_id = aws_cloudfront_distribution.pike.id

  monitoring_subscription {
    realtime_metrics_subscription_config {
      realtime_metrics_subscription_status = "Enabled"
    }
  }
}
