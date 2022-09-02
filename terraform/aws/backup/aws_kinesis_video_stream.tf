resource "aws_kinesis_video_stream" "pike" {
  name                    = "pike-terraform-kinesis-video-stream"
  data_retention_in_hours = 1
  device_name             = "kinesis-video-device-name"
  media_type              = "video/h264"

  kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  tags = {
    pike = "permissions"
    Name = "terraform-kinesis-video-stream"
  }
}
