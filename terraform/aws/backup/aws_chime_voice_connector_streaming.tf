resource "aws_chime_voice_connector_streaming" "pike" {
  disabled                       = false
  voice_connector_id             = aws_chime_voice_connector.pike.id
  data_retention                 = 7
  streaming_notification_targets = ["SQS"]

}
