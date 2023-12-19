resource "aws_vpclattice_service_network_vpc_association" "pike" {
  vpc_identifier             = "vpc-0c33dc8cd64f408c4"
  service_network_identifier = aws_vpclattice_service_network.pike.id
  security_group_ids         = ["sg-05b27cb61c9c46bd2"]
}
