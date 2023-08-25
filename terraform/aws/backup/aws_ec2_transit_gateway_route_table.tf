resource "aws_ec2_transit_gateway_route_table" "pike" {
  transit_gateway_id = aws_ec2_transit_gateway.example.id
}
