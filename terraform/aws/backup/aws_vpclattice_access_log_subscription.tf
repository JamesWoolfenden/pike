resource "aws_vpclattice_access_log_subscription" "pike" {
  resource_identifier = aws_vpclattice_service_network.pike.id
  destination_arn     = aws_s3_bucket.bucket.arn
}

resource "aws_s3_bucket" "bucket" {}
