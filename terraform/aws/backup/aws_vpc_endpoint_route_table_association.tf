resource "aws_vpc_endpoint_route_table_association" "pike" {
  route_table_id  = aws_route_table.example.id
  vpc_endpoint_id = aws_vpc_endpoint.example.id
}
