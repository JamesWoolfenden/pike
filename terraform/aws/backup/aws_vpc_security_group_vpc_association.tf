resource "aws_vpc_security_group_vpc_association" "pike" {
  vpc_id            = aws_vpc.example.id
  security_group_id = aws_security_group.example2.id
}

resource "aws_security_group" "example2" {}
