resource "aws_oam_sink" "pike" {
  name = "ExampleSink"

  tags = {
    Env = "prod"
  }
}
