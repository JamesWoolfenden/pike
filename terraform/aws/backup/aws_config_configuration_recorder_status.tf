resource "aws_config_configuration_recorder_status" "pike" {
  is_enabled = false
  name       = aws_config_configuration_recorder.pike.name
  depends_on = [aws_config_delivery_channel.pike]
}
