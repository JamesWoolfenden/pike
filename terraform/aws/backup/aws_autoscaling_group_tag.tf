resource "aws_autoscaling_group_tag" "pike" {

  tag {
    key                 = "pike"
    value               = "permissions"
    propagate_at_launch = true
  }
  autoscaling_group_name = aws_autoscaling_group.pike.name
}
