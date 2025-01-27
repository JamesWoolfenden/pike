resource "aws_cloudwatch_log_delivery" "example" {
  delivery_source_name     = aws_cloudwatch_log_delivery_source.example.name
  delivery_destination_arn = aws_cloudwatch_log_delivery_destination.example.arn

  field_delimiter = ","

  record_fields = ["event_timestamp", "event"]
}
