data "aws_ec2_service_link_virtual_interface" "pike" {
  id = "slvif-01234567890abcdef"
}

output "aws_ec2_service_link_virtual_interface" {
  value = data.aws_ec2_service_link_virtual_interface.pike
}
