resource "aws_ec2_transit_gateway_multicast_domain_association" "pike" {
  subnet_id                           = aws_subnet.pike.id
  transit_gateway_attachment_id       = aws_ec2_transit_gateway_vpc_attachment.pike.id
  transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain.pike.id
}
