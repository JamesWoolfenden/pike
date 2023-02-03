resource "aws_cloudwatch_metric_stream" "pike" {
  name          = "my-metric-stream"
  role_arn      = "arn:aws:iam::680235478471:role/firehose"
  firehose_arn  = "arn:aws:firehose:eu-west-2:680235478471:deliverystream/PUT-S3-CXcb4"
  output_format = "opentelemetry0.7"

  include_filter {
    namespace = "AWS/EC2"
  }

  include_filter {
    namespace = "AWS/EBS"
  }
  tags = {
    pike = "permission"
    #    another="tag"
  }
}
