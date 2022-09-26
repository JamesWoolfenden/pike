resource "aws_vpn_gateway_attachment" "pike" {
  vpc_id         = "vpc-0c33dc8cd64f408c4"
  vpn_gateway_id = aws_vpn_gateway.vpn.id
}


resource "aws_vpn_gateway" "vpn" {}
