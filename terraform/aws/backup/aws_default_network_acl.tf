resource "aws_default_network_acl" "pike" {
  default_network_acl_id = "acl-01c8f6820c190c9dd"
  ingress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  egress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }
}

output "default_acl" {
  value = aws_default_network_acl.pike
}
