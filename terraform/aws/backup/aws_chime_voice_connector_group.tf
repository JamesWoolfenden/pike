resource "aws_chime_voice_connector_group" "pike" {
  name = "test-group"

  connector {
    voice_connector_id = aws_chime_voice_connector.pike.id
    priority           = 1
  }

  connector {
    voice_connector_id = aws_chime_voice_connector.vc2.id
    priority           = 3
  }

}
