resource "aws_ec2_transit_gateway_vpc_attachment" "pike" {
  subnet_ids         = [aws_subnet.example.id]
  transit_gateway_id = aws_ec2_transit_gateway.example.id
  vpc_id             = aws_vpc.example.id
}
