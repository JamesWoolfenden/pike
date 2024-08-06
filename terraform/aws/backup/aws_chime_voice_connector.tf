resource "aws_chime_voice_connector" "pike" {
  provider           = aws.central
  name               = "connector-test-1"
  require_encryption = true
  aws_region         = "us-east-1"
  tags = {
    pike = "permissions"
  }
}

resource "aws_chime_voice_connector" "vc2" {
  provider           = aws.uswest2
  name               = "connector-test-2"
  require_encryption = true
  aws_region         = "us-west-2"
}
