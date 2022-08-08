resource "aws_cloudwatch_log_group" "name" {
  name = "Yada"
  // kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  retention_in_days = 7
  tags = {
    Environment = "production"
    Application = "serviceA"
  }
}
