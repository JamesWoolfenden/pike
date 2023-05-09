resource "aws_autoscaling_policy" "example" {
  name                   = "example"
  scaling_adjustment     = 4
  adjustment_type        = "ChangeInCapacity"
  cooldown               = 300
  autoscaling_group_name = "group_name"
}
