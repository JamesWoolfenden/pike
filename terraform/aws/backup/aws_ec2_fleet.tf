resource "aws_ec2_fleet" "pike" {
  launch_template_config {
    launch_template_specification {
      launch_template_id = aws_launch_template.example.id
      version            = aws_launch_template.example.latest_version
    }
  }

  target_capacity_specification {
    default_target_capacity_type = "spot"
    total_target_capacity        = 5
  }
}

resource "aws_launch_template" "example" {
  name          = "pike"
  image_id      = "ami-078a289ddf4b09ae0"
  instance_type = "t2.micro"

}
