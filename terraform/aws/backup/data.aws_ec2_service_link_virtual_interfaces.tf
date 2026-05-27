data "aws_ec2_service_link_virtual_interfaces" "pike" {
}

output "aws_ec2_service_link_virtual_interfaces" {
  value = data.aws_ec2_service_link_virtual_interfaces.pike
}
