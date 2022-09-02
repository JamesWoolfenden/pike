resource "aws_volume_attachment" "example" {
  device_name = "/dev/sdh"
  volume_id   = "vol-0227445f299ad9dbb"
  instance_id = "i-0134772b834b9f299"
}
