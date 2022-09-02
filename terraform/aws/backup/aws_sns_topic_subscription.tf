resource "aws_sns_topic_subscription" "user_updates_sqs_target" {
  topic_arn = "arn:aws:sns:eu-west-2:680235478471:processedmessage"
  protocol  = "sqs"
  endpoint  = "arn:aws:sqs:eu-west-2:680235478471:examplequeue"
}
