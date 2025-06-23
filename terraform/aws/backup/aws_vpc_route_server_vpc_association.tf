resource "aws_vpc_route_server_vpc_association" "pike" {
  provider        = aws.central
  route_server_id = aws_vpc_route_server.pike.route_server_id
  vpc_id          = aws_vpc.main.id
}
