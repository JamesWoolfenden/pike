resource "aws_security_group_rule" "rule" {
  from_port         = 81
  protocol          = "tcp"
  security_group_id = "sg-06db45d8099f7f549"
  to_port           = 81
  type              = "ingress"
  cidr_blocks       = ["10.0.0.0/16"]
}
