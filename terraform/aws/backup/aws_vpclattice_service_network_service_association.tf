resource "aws_vpclattice_service_network_service_association" "pike" {
  service_identifier         = aws_vpclattice_service.pike.id
  service_network_identifier = aws_vpclattice_service_network.pike.id
}
