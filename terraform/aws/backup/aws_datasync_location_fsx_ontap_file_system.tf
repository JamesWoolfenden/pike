resource "aws_datasync_location_fsx_ontap_file_system" "pike" {
  fsx_filesystem_arn          = aws_fsx_ontap_file_system.test.arn
  security_group_arns         = [aws_security_group.example.arn]
  storage_virtual_machine_arn = aws_fsx_ontap_storage_virtual_machine.test.arn

  protocol {
    nfs {
      mount_options {
        version = "NFS3"
      }
    }
  }
  tags = {
    pike = "permissions"
  }
}
