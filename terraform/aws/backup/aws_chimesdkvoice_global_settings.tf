resource "aws_chimesdkvoice_global_settings" "pike" {
  provider = aws.central
  voice_connector {
    cdr_bucket = "pike-680235478471"
  }

}
