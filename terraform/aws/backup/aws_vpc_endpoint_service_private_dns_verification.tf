resource "aws_vpc_endpoint_service_private_dns_verification" "pike" {
  service_id = data.aws_vpc_endpoint_service.example.service_id
}

data "aws_vpc_endpoint_service" "example" {
  service      = "s3"
  service_type = "Gateway"
}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_vpc_endpoint" "example" {
  service_name = data.aws_vpc_endpoint_service.example.service_name
  vpc_id       = aws_vpc.example.id
}
