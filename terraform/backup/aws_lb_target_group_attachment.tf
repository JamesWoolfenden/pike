
resource "aws_lb_target_group_attachment" "test" {
  target_group_arn = aws_lb_target_group.test.arn
  target_id        = "i-096e371e3ce592fc0"
  port             = 80
}
