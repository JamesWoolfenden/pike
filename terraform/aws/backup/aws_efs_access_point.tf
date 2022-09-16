resource "aws_efs_access_point" "pike" {
  file_system_id = "fs-0898bc857b16b617a"
  tags = {
    pike = "permissions"
    //   delete="me"
  }
}
