data "aws_efs_file_system" "pike" {
  file_system_id = aws_efs_file_system.pike.id
}

data "aws_efs_file_system" "pike-tag" {
  tags = {
    pike = "permissions"
  }
}
