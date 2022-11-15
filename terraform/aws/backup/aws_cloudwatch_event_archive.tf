resource "aws_cloudwatch_event_archive" "pike" {
  name             = "order-archive"
  event_source_arn = aws_cloudwatch_event_bus.pike.arn
}
