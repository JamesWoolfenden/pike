# resource "aws_vpc_route_server_endpoint" "pike" {
#   provider = aws.central
#
#   route_server_id = aws_vpc_route_server.pike.route_server_id
#   subnet_id       = aws_subnet.first.id
#
#   tags = {
#     Name = "Endpoint A"
#   }
# }


resource "aws_vpc" "main" {
  provider = aws.central

  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "first" {
  provider = aws.central

  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"
}
