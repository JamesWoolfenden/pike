resource "aws_secretsmanager_secret" "example" {
  description = "My secret"
  kms_key_id  = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  name        = "pike"
  //policy = {}
  recovery_window_in_days = 0
  tags = {
    "pike" = "permissions"
  }
}
