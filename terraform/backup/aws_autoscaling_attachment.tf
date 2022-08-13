resource "aws_autoscaling_attachment" "asg_attachment_bar" {
  autoscaling_group_name = "test"
  elb                    = "pike"
}
