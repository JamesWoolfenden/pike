resource "aws_fsx_openzfs_volume" "pike" {
  name             = "pike-child"
  parent_volume_id = aws_fsx_openzfs_file_system.pike.root_volume_id
  tags = {
    pike = "permissions"
  }
}
