resource "aws_dx_hosted_transit_virtual_interface_accepter" "pike" {
  provider             = aws.accepter
  virtual_interface_id = aws_dx_hosted_transit_virtual_interface.creator.id
  dx_gateway_id        = aws_dx_gateway.example.id

  tags = {
    Side = "Accepter"
  }
}