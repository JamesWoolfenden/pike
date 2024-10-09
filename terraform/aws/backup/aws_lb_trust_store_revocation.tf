resource "aws_lb_trust_store_revocation" "pike" {
  trust_store_arn = aws_lb_trust_store.pike.arn

  revocations_s3_bucket = aws_s3_bucket.truststore.bucket
  revocations_s3_key    = "trust"
}


resource "aws_s3_bucket" "truststore" {}
