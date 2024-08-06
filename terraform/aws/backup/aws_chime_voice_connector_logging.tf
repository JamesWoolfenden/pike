resource "aws_chime_voice_connector_logging" "pike" {
  enable_sip_logs          = true
  enable_media_metric_logs = true
  voice_connector_id       = aws_chime_voice_connector.pike.id
}
