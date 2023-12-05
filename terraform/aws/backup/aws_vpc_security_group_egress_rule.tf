resource "aws_vpc_security_group_egress_rule" "pike" {
  security_group_id = aws_security_group.pike.id

  cidr_ipv4   = "10.0.0.0/8"
  from_port   = 80
  ip_protocol = "tcp"
  to_port     = 80
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

resource "aws_security_group" "pike" {}
