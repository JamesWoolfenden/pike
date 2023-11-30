resource "aws_lb_cookie_stickiness_policy" "pike" {
  name                     = "pike"
  load_balancer            = aws_elb.lb.id
  lb_port                  = 80
  cookie_expiration_period = 600
}

resource "aws_elb" "lb" {
  name               = "test-lb"
  availability_zones = ["eu-west-2a"]

  listener {
    instance_port     = 8000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }
}
