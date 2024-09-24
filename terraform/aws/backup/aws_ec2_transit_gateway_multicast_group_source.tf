resource "aws_ec2_transit_gateway_multicast_group_source" "pike" {
  group_ip_address                    = "224.0.0.1"
  network_interface_id                = aws_network_interface.pike.id
  transit_gateway_multicast_domain_id = aws_ec2_transit_gateway_multicast_domain.pike.id
}
