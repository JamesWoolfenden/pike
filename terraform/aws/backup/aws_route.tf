resource "aws_route" "example" {
  route_table_id         = "rtb-034b3dd8db811ca44"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "igw-035a0222611f0437b"
}
