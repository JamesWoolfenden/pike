resource "aws_default_subnet" "pike" {
  availability_zone = "eu-west-2a"
  tags = {
    Name = "Default Subnet"
    pike = "permissions"
  }
}
