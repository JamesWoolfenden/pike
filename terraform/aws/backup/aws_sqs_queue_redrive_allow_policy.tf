resource "aws_sqs_queue_redrive_allow_policy" "pike" {
  queue_url = "https://sqs.eu-west-2.amazonaws.com/680235478471/examplequeue"

  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = ["arn:aws:sqs:eu-west-2:680235478471:deadletter"]
  })
}
