resource "aws_appstream_fleet" "pike" {
  name          = "NAME"
  image_name    = "Amazon-AppStream2-Sample-Image-03-11-2023"
  instance_type = "stream.standard.small"

  compute_capacity {
    desired_instances = 1
  }
}
