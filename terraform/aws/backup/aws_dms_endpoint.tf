resource "aws_dms_endpoint" "pike" {
  certificate_arn             = aws_dms_certificate.pike.certificate_arn
  database_name               = "test"
  endpoint_id                 = "test-dms-endpoint-tf"
  endpoint_type               = "source"
  engine_name                 = "aurora"
  extra_connection_attributes = ""
  kms_key_arn                 = data.aws_kms_key.example.arn
  password                    = "test"
  port                        = 3306
  server_name                 = "test"
  ssl_mode                    = "none"

  tags = {
    Name = "test"
  }

  username = "test"
}

data "aws_acm_certificate" "pike" {
  domain = "pike-iac.com"
}
