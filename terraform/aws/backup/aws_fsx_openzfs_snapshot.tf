resource "aws_fsx_openzfs_snapshot" "pike" {
  name      = "pike"
  volume_id = aws_fsx_openzfs_file_system.pike.root_volume_id
  tags = {
    pike = "permissions"
  }
}
