data "aws_ec2_serial_console_access" "pike" {}

output "access" {
  value = data.aws_ec2_serial_console_access.pike
}
