resource "aws_opensearchserverless_vpc_endpoint" "pike" {
  provider   = aws.central
  name       = "myendpoint"
  subnet_ids = [aws_subnet.example.id]
  vpc_id     = aws_vpc.example.id
}

resource "aws_subnet" "example" {
  provider   = aws.central
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.0.0/24"
}

resource "aws_vpc" "example" {
  provider             = aws.central
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}
