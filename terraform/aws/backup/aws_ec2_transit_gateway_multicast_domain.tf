resource "aws_ec2_transit_gateway_multicast_domain" "pike" {
  transit_gateway_id = aws_ec2_transit_gateway.pike.id

  static_sources_support = "enable"

  tags = {
    Name = "Transit_Gateway_Multicast_Domain_Example"
  }
}
