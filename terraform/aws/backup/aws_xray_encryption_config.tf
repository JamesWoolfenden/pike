resource "aws_xray_encryption_config" "pike" {
  type   = "KMS"
  key_id = data.aws_kms_key.pike.arn
}

data "aws_kms_key" "pike" {
  key_id = "34cdce9a-2322-427c-91bb-b572f435c032"
}
