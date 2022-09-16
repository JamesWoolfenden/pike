resource "aws_efs_file_system" "pike" {
  availability_zone_name = "eu-west-2a"
  encrypted              = true
  kms_key_id             = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  lifecycle_policy {
    transition_to_ia = "AFTER_14_DAYS"
  }
  performance_mode = "generalPurpose"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
