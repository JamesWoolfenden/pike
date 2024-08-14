resource "aws_cloudfront_realtime_log_config" "pike" {
  name          = "example"
  sampling_rate = 75
  fields        = ["timestamp", "c-ip"]

  endpoint {
    stream_type = "Kinesis"

    kinesis_stream_config {
      role_arn   = aws_iam_role.example.arn
      stream_arn = aws_kinesis_stream.example.arn
    }
  }

}
