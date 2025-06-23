resource "aws_kinesis_stream" "example" {
  name        = "amazon-workspaces-web-user-access-logging-stream"
  shard_count = 1
}

resource "aws_workspacesweb_user_access_logging_settings" "example" {
  kinesis_stream_arn = aws_kinesis_stream.example.arn
  tags = {
    Name        = "example-user-access-logging-settings"
    Environment = "Production"
  }
}
