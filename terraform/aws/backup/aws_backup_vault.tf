resource "aws_backup_vault" "pike" {
  name        = "pike"
  kms_key_arn = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  tags = {
    pike = "Permissions"
  }
}
