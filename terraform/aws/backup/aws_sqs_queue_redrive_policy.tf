resource "aws_sqs_queue_redrive_policy" "pike" {

  queue_url = "https://sqs.eu-west-2.amazonaws.com/680235478471/examplequeue"

  redrive_policy = jsonencode({
    deadLetterTargetArn = ["arn:aws:sqs:eu-west-2:680235478471:deadletter"]
    maxReceiveCount     = 4
  })
}
