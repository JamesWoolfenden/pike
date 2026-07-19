resource "aws_alb_target_group" "test" {
  name     = "tf-example-alb-tg2"
  port     = 80
  protocol = "HTTP"
  vpc_id   = "vpc-0c33dc8cd64f408c4"
  tags = {
    pike = "permission"
  }
}
