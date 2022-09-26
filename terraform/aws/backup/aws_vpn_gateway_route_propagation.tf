resource "aws_vpn_gateway_route_propagation" "pike" {
  route_table_id = aws_route_table.vpn.id
  vpn_gateway_id = aws_vpn_gateway.vpn.id
}

resource "aws_route_table" "vpn" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
}

resource "aws_vpn_gateway" "vpn" {}
