resource "aws_autoscaling_group" "pike" {
  availability_zones        = ["eu-west-2a"]
  name                      = "terraform-test-foobar5"
  max_size                  = 1
  min_size                  = 1
  health_check_grace_period = 300
  health_check_type         = "ELB"
  force_delete              = true
  termination_policies      = ["OldestInstance"]
  launch_configuration      = aws_launch_configuration.pike.name
  depends_on = [
    aws_launch_configuration.pike
  ]
  tag {
    key                 = "pike"
    propagate_at_launch = false
    value               = "permissions"
  }
}

resource "aws_autoscaling_schedule" "pike" {
  scheduled_action_name  = "foobar"
  min_size               = 0
  max_size               = 1
  desired_capacity       = 0
  start_time             = "2025-12-11T18:00:00Z"
  end_time               = "2025-12-12T06:00:00Z"
  autoscaling_group_name = aws_autoscaling_group.pike.name
}


resource "aws_launch_configuration" "pike" {
  name              = "pike2"
  image_id          = "ami-09a2a0f7d2db8baca"
  instance_type     = "t3.micro"
  security_groups   = ["sg-06b8c96aaccf3a2a1"]
  enable_monitoring = false
  metadata_options {
    http_tokens = "required"
  }
}
