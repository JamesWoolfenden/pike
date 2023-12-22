resource "aws_autoscaling_traffic_source_attachment" "pike" {
  autoscaling_group_name = aws_autoscaling_group.pike.id

  traffic_source {
    identifier = aws_lb_target_group.pike.arn
    type       = "elbv2"
  }
}

resource "aws_lb_target_group" "pike" {
  name     = "tf-example-lb-tg3"
  port     = 80
  protocol = "HTTP"
  vpc_id   = "vpc-00ea5eff890b0e212"
  tags = {
    pike = "permission"
    #    tag="deletable"
  }
}
