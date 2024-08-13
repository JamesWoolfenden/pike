resource "aws_vpn_connection_route" "pike" {
  destination_cidr_block = "192.168.10.0/ 24"
  vpn_connection_id      = aws_vpn_connection.main.id
}
