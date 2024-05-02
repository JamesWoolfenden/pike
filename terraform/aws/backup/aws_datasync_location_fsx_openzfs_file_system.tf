resource "aws_datasync_location_fsx_openzfs_file_system" "pike" {

  fsx_filesystem_arn  = aws_fsx_openzfs_file_system.example.arn
  security_group_arns = [aws_security_group.example.arn]

  protocol {
    nfs {
      mount_options {
        version = "AUTOMATIC"
      }
    }
  }

  tags = {
    pike = "permissions"
  }
}

resource "aws_fsx_openzfs_file_system" "example" {
  storage_capacity    = 64
  subnet_ids          = ["subnet-09ff91b5b0adb1fd4"]
  deployment_type     = "SINGLE_AZ_1"
  throughput_capacity = 64
  kms_key_id          = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  security_group_ids  = ["sg-06b8c96aaccf3a2a1"]
  tags = {
    pike = "permissions"
  }
}
