resource "aws_lightsail_instance_public_ports" "pike" {
  instance_name = aws_lightsail_instance.pike.name

  port_info {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
  }

}
