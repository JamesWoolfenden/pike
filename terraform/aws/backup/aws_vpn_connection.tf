resource "aws_vpn_connection" "example" {
  customer_gateway_id                     = "customer_gateway"
  outside_ip_address_type                 = "PrivateIpv4"
  transit_gateway_id                      = "transit_gateway"
  transport_transit_gateway_attachment_id = "transit_gateway_attachment"
  type                                    = "ipsec.1"

  tags = {
    Name = "terraform_ipsec_vpn_example"
  }
}
