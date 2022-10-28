resource "aws_lambda_function_event_invoke_config" "pike" {
  function_name = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"

  destination_config {
    on_failure {
      destination = "arn:aws:sqs:eu-west-2:680235478471:deadletter"
    }

    on_success {
      destination = "arn:aws:sns:eu-west-2:680235478471:pike"
    }
  }
}
