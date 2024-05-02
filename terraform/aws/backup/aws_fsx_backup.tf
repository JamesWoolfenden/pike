resource "aws_fsx_backup" "pike" {
  file_system_id = aws_fsx_lustre_file_system.pike.id
  tags = {
    pike = "permissions"
  }
}
