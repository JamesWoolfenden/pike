resource "aws_ec2_client_vpn_authorization_rule" "pike" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.pike.id
  target_network_cidr    = aws_subnet.example.cidr_block
  authorize_all_groups   = true
}
