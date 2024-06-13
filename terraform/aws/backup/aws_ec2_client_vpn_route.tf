resource "aws_ec2_client_vpn_route" "pike" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.pike.id
  destination_cidr_block = "0.0.0.0/0"
  target_vpc_subnet_id   = aws_ec2_client_vpn_network_association.pike.subnet_id
}
