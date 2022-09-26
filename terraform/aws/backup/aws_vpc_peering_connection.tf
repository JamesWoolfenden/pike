resource "aws_vpc_peering_connection" "pike" {
  peer_vpc_id = aws_vpc.peer.id
  vpc_id      = aws_vpc.base.id

  tags = {
    pike = "permissions"
  }
}

resource "aws_vpc" "peer" {
  cidr_block = "10.0.0.0/24"
}

resource "aws_vpc" "base" {
  cidr_block = "10.1.0.0/24"
}
