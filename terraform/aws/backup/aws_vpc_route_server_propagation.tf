resource "aws_vpc_route_server_propagation" "example" {
  provider        = aws.central
  route_server_id = aws_vpc_route_server.pike.route_server_id
  route_table_id  = aws_route_table.pike.id
  depends_on = [
  aws_vpc_route_server_vpc_association.pike]
}


resource "aws_route_table" "pike" {
  provider = aws.central
  vpc_id   = aws_vpc.main.id
}
