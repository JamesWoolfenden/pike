resource "aws_vpc_dhcp_options_association" "pike" {
  dhcp_options_id = aws_vpc_dhcp_options.pike.id
  vpc_id          = "vpc-0c33dc8cd64f408c4"
}
