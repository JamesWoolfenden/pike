#resource "aws_sqs_queue" "user_updates_queue" {
#  name = "user-updates.fifo"
#  fifo_queue = true
##  sqs_managed_sse_enabled=true
##  sqs/SetQueueAttributes
##  kms_master_key_id                 = "alias/aws/sqs"
##  kms_data_key_reuse_period_seconds = 300
#  visibility_timeout_seconds=3600
#
#  policy = <<POLICY
#{
#  "Id": "Policy1660312635910",
#  "Version": "2012-10-17",
#  "Statement": [
#    {
#      "Sid": "Stmt1660312634684",
#      "Action": "sqs:*",
#      "Effect": "Allow",
#      "Resource": "arn:aws:sqs:eu-west-2:680235478471:user-updates.fifo",
#      "Principal": "*"
#    }
#  ]
#}
#POLICY
#
#  tags = {
#    pike="permissions"
#  }
#}
#
#resource "aws_sqs_queue" "q" {
#  name = "examplequeue"
#}

resource "aws_sqs_queue_policy" "test" {
  queue_url = "https://sqs.eu-west-2.amazonaws.com/680235478471/examplequeue"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "arn:aws:sqs:eu-west-2:680235478471:examplequeue"
    }
  ]
}
POLICY
}
