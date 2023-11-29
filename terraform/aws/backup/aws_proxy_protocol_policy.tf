resource "aws_elb" "lb" {
  name = "test-lb"

  availability_zones = ["us-east-1a"]

  listener {
    instance_port     = 25
    instance_protocol = "tcp"
    lb_port           = 25
    lb_protocol       = "tcp"
  }

  listener {
    instance_port     = 587
    instance_protocol = "tcp"
    lb_port           = 587
    lb_protocol       = "tcp"
  }
}

resource "aws_proxy_protocol_policy" "smtp" {
  load_balancer  = aws_elb.lb.name
  instance_ports = ["25", "587"]
}
