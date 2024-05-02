resource "aws_datasync_location_fsx_lustre_file_system" "pike" {
  fsx_filesystem_arn  = aws_fsx_lustre_file_system.pike.arn
  security_group_arns = [aws_security_group.example.arn]

  tags = {
    pike = "permissions"
  }
}

resource "aws_fsx_lustre_file_system" "pike" {
  subnet_ids = [data.aws_subnets.example.ids[0]]
}
