resource "aws_vpc_peering_connection_options" "pike" {
  vpc_peering_connection_id = aws_vpc_peering_connection.pike.id
}
