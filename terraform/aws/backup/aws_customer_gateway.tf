resource "aws_customer_gateway" "pike" {
  bgp_asn    = 65000
  ip_address = "172.83.124.10"
  type       = "ipsec.1"

  tags = {
    pike = "permissions"
    Name = "main-customer-gateway"
  }
}
