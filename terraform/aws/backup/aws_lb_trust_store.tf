resource "aws_lb_trust_store" "pike" {
  ca_certificates_bundle_s3_bucket = aws_s3_bucket.truststore.bucket
  ca_certificates_bundle_s3_key    = "trust"
}
