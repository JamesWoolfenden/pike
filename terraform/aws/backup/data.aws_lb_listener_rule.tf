data "aws_lb_listener_rule" "pike" {
}

output "aws_lb_listener_rule" {
  value = data.aws_lb_listener_rule.pike
}
