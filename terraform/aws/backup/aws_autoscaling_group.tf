resource "aws_autoscaling_group" "example" {
  launch_configuration = aws_launch_configuration.test.name
  max_size             = 0
  min_size             = 0
  name                 = "test"
  vpc_zone_identifier  = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7", "subnet-05a6a6de2f4989d22"]
  tag {
    key                 = "Key"
    value               = "Value"
    propagate_at_launch = true
  }
}

resource "aws_launch_configuration" "test" {
  image_id          = "ami-09a2a0f7d2db8baca"
  instance_type     = "t3.micro"
  security_groups   = ["sg-0d2f425bb20ccbd04"]
  enable_monitoring = false
}
