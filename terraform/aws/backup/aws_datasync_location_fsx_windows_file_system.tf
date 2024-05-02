resource "aws_datasync_location_fsx_windows_file_system" "pike" {
  fsx_filesystem_arn  = aws_fsx_windows_file_system.pike.arn
  user                = "SomeUser"
  password            = "SuperSecretPassw0rd"
  security_group_arns = [aws_security_group.example.arn]

  tags = {
    pike = "permissions"
  }
}

resource "aws_security_group" "example" {}

resource "aws_fsx_windows_file_system" "pike" {
  subnet_ids          = data.aws_subnets.example.ids
  throughput_capacity = 8
}

data "aws_subnets" "example" {}
