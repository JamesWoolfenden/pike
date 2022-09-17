resource "aws_sfn_state_machine" "pike" {
  name     = "pike"
  role_arn = "arn:aws:iam::680235478471:role/sfn"

  definition = <<EOF
{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function.",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:eu-west-2:680235478471:function:message-processor",
      "End": true
    }
  }
}
EOF

  logging_configuration {
    log_destination        = "arn:aws:logs:eu-west-2:680235478471:log-group:/aws/lambda/message-processor:*"
    include_execution_data = true
    level                  = "ERROR"
  }

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
