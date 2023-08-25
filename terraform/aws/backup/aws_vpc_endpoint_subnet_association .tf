resource "aws_vpc_endpoint_subnet_association" "pike" {
  vpc_endpoint_id = aws_vpc_endpoint.example.id
  subnet_id       = aws_subnet.example.id
}
