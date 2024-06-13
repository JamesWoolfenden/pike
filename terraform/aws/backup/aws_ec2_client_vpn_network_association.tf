resource "aws_ec2_client_vpn_network_association" "pike" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.pike.id
  subnet_id              = aws_subnet.example.id
}
resource "aws_subnet" "example" {
  vpc_id     = "vpc-06074a092930bc809" #vpc-0c9622709bb598517"
  cidr_block = "10.0.1.0/24"

  #   tags = {
  #     Name = "Main"
  #   }
}
