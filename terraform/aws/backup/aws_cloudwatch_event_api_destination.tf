resource "aws_cloudwatch_event_api_destination" "pike" {
  name                             = "api-destination"
  description                      = "An API Destination - update"
  invocation_endpoint              = "https://httpbin.org/get"
  http_method                      = "POST"
  invocation_rate_limit_per_second = 20
  connection_arn                   = aws_cloudwatch_event_connection.pike.arn
}
