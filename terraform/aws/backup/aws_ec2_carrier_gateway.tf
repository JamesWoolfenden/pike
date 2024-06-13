resource "aws_ec2_carrier_gateway" "pike" {
  vpc_id = "ami-078a289ddf4b09ae0"
  tags = {
    pike = "permissions"
  }
}
