resource "aws_kms_external_key" "pike" {
  provider = aws.central

  description             = "Multi-Region primary key"
  deletion_window_in_days = 30
  multi_region            = true
  enabled                 = true

}
