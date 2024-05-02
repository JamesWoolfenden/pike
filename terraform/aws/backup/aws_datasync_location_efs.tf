resource "aws_datasync_location_efs" "pike" {
  #   The below example uses aws_efs_mount_target as a reference to ensure a mount target already exists when resource creation occurs.
  #   You can accomplish the same behavior with depends_on or an aws_efs_mount_target data source reference.
  efs_file_system_arn = aws_efs_mount_target.pike.file_system_arn

  ec2_config {
    security_group_arns = [aws_security_group.example.arn]
    subnet_arn          = data.aws_subnet.selected.arn
  }

  tags = {
    pike = "permissions"
  }
}

data "aws_subnet" "selected" {
  id = data.aws_subnets.example.ids[0]
}

resource "aws_efs_mount_target" "pike" {}
