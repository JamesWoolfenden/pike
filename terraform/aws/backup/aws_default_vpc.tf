resource "aws_default_vpc" "pike" {

  tags = {
    Name = "Default VPC"
    pike = "permissions"
    //  delete="me"
  }
}


output "default_vpc" {
  value = aws_default_vpc.pike
}
