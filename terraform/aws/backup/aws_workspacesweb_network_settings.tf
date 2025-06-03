
resource "aws_vpc" "example" {
  provider   = aws.central
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "example" {
  provider = aws.central
  count    = 2

  vpc_id            = aws_vpc.example.id
  cidr_block        = cidrsubnet(aws_vpc.example.cidr_block, 8, count.index)
  availability_zone = data.aws_availability_zones.available.names[count.index]
}

resource "aws_security_group" "example1" {
  provider = aws.central
  count    = 2

  vpc_id = aws_vpc.example.id
  name   = "example-sg-${count.index}$"
}

resource "aws_workspacesweb_network_settings" "example" {
  provider           = aws.central
  vpc_id             = aws_vpc.example.id
  subnet_ids         = [aws_subnet.example[0].id, aws_subnet.example[1].id]
  security_group_ids = [aws_security_group.example[0].id, aws_security_group.example[1].id]
}

data "aws_availability_zones" "available" {
  provider = aws.central
}


resource "aws_security_group" "example" {
  provider = aws.central
  vpc_id   = aws_vpc.example.id
  count    = 2
}
