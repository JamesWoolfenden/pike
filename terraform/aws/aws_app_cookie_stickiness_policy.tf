resource "aws_app_cookie_stickiness_policy" "pike" {
  name          = "foobarpolicy"
  load_balancer = aws_elb.lb.name
  lb_port       = 80
  cookie_name   = "MyAppCookie"
}
