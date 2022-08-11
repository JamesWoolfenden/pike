resource "aws_default_security_group" "example" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
  ingress {
    protocol    = -1
    self        = true
    from_port   = 0
    to_port     = 0
    description = "ingress"
  }

  egress {
    description = "egress"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    pike = "permissions"
    // another="permission"
  }
}
