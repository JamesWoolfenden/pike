

resource "aws_instance" "instance" {
  instance_type = "t2.micro"
  ami           = "ami-091f18e98bc129c4e"

  tags = {
    type = "terraform-test-instance"
  }
}

resource "aws_security_group" "sg" {
  tags = {
    type = "terraform-test-security-group"
  }
}

resource "aws_network_interface_sg_attachment" "sg_attachment" {
  security_group_id    = aws_security_group.sg.id
  network_interface_id = aws_instance.instance.primary_network_interface_id
}
