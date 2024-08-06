resource "aws_chime_voice_connector_termination_credentials" "pike" {
  voice_connector_id = aws_chime_voice_connector.pike.id

  credentials {
    username = "test"
    password = "test!"
  }

  depends_on = [aws_chime_voice_connector_termination.pike]
}
