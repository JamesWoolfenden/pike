resource "aws_launch_configuration" "pike" {
  image_id          = "ami-09a2a0f7d2db8baca"
  instance_type     = "t3.micro"
  security_groups   = ["sg-0d2f425bb20ccbd04"]
  enable_monitoring = false
}
