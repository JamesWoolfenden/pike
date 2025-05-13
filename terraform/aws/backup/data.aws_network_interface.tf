data "aws_network_interface" "pike" {
}

output "aws_network_interface" {
  value = data.aws_network_interface.pike
}
