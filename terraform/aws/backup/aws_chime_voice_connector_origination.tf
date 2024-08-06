resource "aws_chime_voice_connector_origination" "pike" {
  disabled           = false
  voice_connector_id = aws_chime_voice_connector.pike.id

  route {
    host     = "127.0.0.1"
    port     = 8081
    protocol = "TCP"
    priority = 1
    weight   = 1
  }

  route {
    host     = "127.0.0.2"
    port     = 8082
    protocol = "TCP"
    priority = 2
    weight   = 10
  }
}
