resource "aws_alb_listener_rule" "pike" {
  listener_arn = aws_alb_listener.example.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.test.arn
  }

  condition {
    path_pattern {
      values = ["/static/*"]
    }
  }

  condition {
    host_header {
      values = ["example.com"]
    }
  }
}
