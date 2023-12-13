resource "aws_cloudwatch_log_destination" "pike" {
  name       = "pike"
  role_arn   = "arn:aws:iam::680235478471:role/CWLtoKinesisRole"
  target_arn = aws_kinesis_stream.pike.arn
  tags = {
    pike = "permissions"
  }
}

resource "aws_kinesis_stream" "pike" {
  name             = "pike"
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  stream_mode_details {
    stream_mode = "PROVISIONED"
  }

  tags = {
    Environment = "test"
  }
}

resource "aws_kinesis_stream_consumer" "pike" {
  name       = "pike"
  stream_arn = aws_kinesis_stream.pike.arn
}
