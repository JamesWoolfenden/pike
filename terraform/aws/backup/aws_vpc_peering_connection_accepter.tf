resource "aws_vpc_peering_connection_accepter" "pike" {
  vpc_peering_connection_id = aws_vpc_peering_connection.pike.id
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
