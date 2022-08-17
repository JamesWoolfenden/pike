resource "aws_vpn_gateway" "example" {
  vpc_id = "vpc-03036aea287d9ee9b"

  tags = {
    pike = "permissions"
  }
}
