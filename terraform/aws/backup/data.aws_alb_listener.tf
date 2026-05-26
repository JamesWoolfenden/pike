data "aws_alb_listener" "pike" {
}

output "aws_alb_listener" {
  value = data.aws_alb_listener.pike
}
