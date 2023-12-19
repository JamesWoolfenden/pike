resource "aws_kms_replica_key" "pike" {
  description             = "Multi-Region replica key"
  deletion_window_in_days = 7
  primary_key_arn         = aws_kms_key.primary.arn
}

resource "aws_kms_key" "primary" {
  provider = aws.central

  description             = "Multi-Region primary key"
  deletion_window_in_days = 30
  multi_region            = true
}
