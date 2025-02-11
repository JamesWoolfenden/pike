
resource "aws_vpc_endpoint_security_group_association" "sg_ec2" {
  vpc_endpoint_id   = aws_vpc_endpoint.example.id
  security_group_id = aws_security_group.sg.id
}


resource "aws_security_group" "sg" {}
