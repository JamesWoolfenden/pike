data "aws_ec2_transit_gateway_dx_gateway_attachment" "pike" {
  transit_gateway_id = aws_ec2_transit_gateway.pike.id
  dx_gateway_id      = aws_dx_gateway.pike.id
}
