resource "aws_dx_gateway_association" "pike" {
  dx_gateway_id         = aws_dx_gateway.example.id
  associated_gateway_id = aws_ec2_transit_gateway.example.id

  allowed_prefixes = [
    "10.255.255.0/30",
    "10.255.255.8/30",
  ]
}