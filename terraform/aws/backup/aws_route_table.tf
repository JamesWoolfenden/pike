resource "aws_route_table" "example" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
  route  = []

  tags = {
    Name = "example"
  }
}
