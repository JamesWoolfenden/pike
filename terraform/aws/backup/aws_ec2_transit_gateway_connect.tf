resource "aws_ec2_transit_gateway_connect" "pike" {
  transport_attachment_id = aws_ec2_transit_gateway_vpc_attachment.pike.id
  transit_gateway_id      = aws_ec2_transit_gateway.pike.id
}
