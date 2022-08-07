resource "aws_security_group" "broker" {
  description = "Managed by Terraform"
  egress {
    #tfsec:ignore:AWS009
    cidr_blocks = ["0.0.0.0/0"]
    description = "Outbound"
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
  }

  ingress {
    cidr_blocks = var.ingress
    description = "MQ port"
    from_port   = 61616
    protocol    = "tcp"
    self        = false
    to_port     = 61616
  }


  name   = var.security_group_name
  vpc_id = var.vpc_id
  tags   = var.common_tags
}
