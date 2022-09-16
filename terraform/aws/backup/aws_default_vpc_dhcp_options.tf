resource "aws_default_vpc_dhcp_options" "pike" {
  tags = {
    Name = "Default DHCP Option Set"
    pike = "permissions"
  }
}
