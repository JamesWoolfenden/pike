resource "aws_lambda_function_url" "pike" {
  authorization_type = "NONE"
  function_name      = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
}

output "lambda_endpoint" {
  value = aws_lambda_function_url.pike
}
