resource "aws_efs_replication_configuration" "pike" {
  source_file_system_id = "fs-0898bc857b16b617a"
  destination {
    availability_zone_name = "eu-west-1a"
    //kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
    region = "eu-west-1"
  }
}

output "replication" {
  value = aws_efs_replication_configuration.pike
}
