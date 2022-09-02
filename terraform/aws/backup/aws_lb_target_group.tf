resource "aws_lb_target_group" "test" {
  name     = "tf-example-lb-tg2"
  port     = 80
  protocol = "HTTP"
  vpc_id   = "vpc-0c33dc8cd64f408c4"
  tags = {
    pike = "permission"
    #    tag="deletable"
  }
}
