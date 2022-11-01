resource "aws_s3_bucket_notification" "pike" {
  bucket = "snsjgwtest"

  // eventbridge = true

  #  lambda_function {
  #    lambda_function_arn = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
  #    events              = ["s3:ObjectCreated:*"]
  #    filter_prefix       = "AWSLogs/"
  #    filter_suffix       = ".log"
  #  }
  topic {
    topic_arn     = aws_sns_topic.topic.arn
    events        = ["s3:ObjectCreated:*"]
    filter_suffix = ".log"
  }
  #
  #  queue {
  #    queue_arn     = "arn:aws:sqs:eu-west-2:680235478471:examplequeue"
  #    events        = ["s3:ObjectCreated:*"]
  #    filter_suffix = ".log"
  #  }
}

#resource "aws_sns_topic" "topic" {
#  name = "s3-event-notification-topic"
#
#  policy = <<POLICY
#{
#    "Version":"2012-10-17",
#    "Statement":[{
#        "Effect": "Allow",
#        "Principal": { "Service": "s3.amazonaws.com" },
#        "Action": "SNS:Publish",
#        "Resource": "arn:aws:sns:*:*:s3-event-notification-topic",
#        "Condition":{
#            "ArnLike":{"aws:SourceArn":"arn:aws:s3:::snsjgwtest"}
#        }
#    }]
#}
#POLICY
#}
