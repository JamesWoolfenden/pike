resource "aws_network_interface" "example" {
  subnet_id      = "subnet-0243b982356b4a0f0"
  description    = "An interface"
  interface_type = "trunk"
  #  ipv4_prefix_count = 1
  tags = {
    pike = "permissions"
  }
}
