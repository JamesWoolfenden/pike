resource "aws_vpc_endpoint_connection_notification" "pike" {
  vpc_endpoint_service_id     = aws_vpc_endpoint_service.pike.id
  connection_notification_arn = aws_sns_topic.pike.arn
  connection_events           = ["Accept", "Reject"]
}
