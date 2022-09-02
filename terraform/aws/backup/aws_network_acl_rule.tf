resource "aws_network_acl_rule" "pike" {
  network_acl_id = "acl-0260ab016a368c5a9"
  rule_number    = 200
  egress         = false
  protocol       = "tcp"
  rule_action    = "allow"
  cidr_block     = "10.0.0.0/16"
  from_port      = 22
  to_port        = 22
}
