resource "aws_alb_listener" "example" {
  load_balancer_arn = aws_alb.test.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.test.arn
  }
}
