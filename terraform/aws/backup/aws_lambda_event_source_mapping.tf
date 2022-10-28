resource "aws_lambda_event_source_mapping" "pike" {
  function_name    = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
  enabled          = true
  event_source_arn = "arn:aws:sqs:eu-west-2:680235478471:deadletter"
}
