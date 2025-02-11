
resource "aws_vpc_endpoint_private_dns" "pike" {
  vpc_endpoint_id     = aws_vpc_endpoint.example.id
  private_dns_enabled = true
}

resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

resource "aws_vpc_endpoint" "example" {

  vpc_id            = aws_vpc.example.id
  service_name      = "com.amazonaws.eu-west-2.ec2"
  vpc_endpoint_type = "Interface"

  security_group_ids = [
    aws_security_group.sg1.id,
  ]

  private_dns_enabled = true
}

resource "aws_security_group" "sg1" {}
