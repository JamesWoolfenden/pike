resource "aws_vpc_dhcp_options" "pike" {
  domain_name_servers = ["8.8.8.8", "8.8.4.4"]
  tags = {
    pike = "permissions"
  }
}
