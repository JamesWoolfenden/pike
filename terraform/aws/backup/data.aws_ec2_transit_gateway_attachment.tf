data "aws_ec2_transit_gateway_attachment" "pike" {
  filter {
    name   = "transit-gateway-id"
    values = [aws_ec2_transit_gateway.pike.id]
  }

  filter {
    name   = "resource-type"
    values = ["peering"]
  }
}
