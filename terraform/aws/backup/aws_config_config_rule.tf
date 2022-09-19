resource "aws_config_config_rule" "pike" {
  name        = "pike"
  description = "Needs description"
  source {
    owner             = "AWS"
    source_identifier = "S3_BUCKET_VERSIONING_ENABLED"
  }

  tags = {
    pike = "permissions"
    //delete = "me"
  }

  depends_on = [aws_config_configuration_recorder.pike]
}
