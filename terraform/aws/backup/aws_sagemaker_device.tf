resource "aws_sagemaker_device" "pike" {
  device_fleet_name = aws_sagemaker_device_fleet.pike.device_fleet_name

  device {
    device_name = "example"
  }
}
