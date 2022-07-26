resource "aws_lambda_function" "examplea" {
  function_name = "examplea"
  role          = "arn:aws:iam::680235478471:role/lambda_basic"
  handler       = "anyoldguff"
  runtime       = "python3.8"
  filename      = "todo.zip"
  # tags = {
  #   createdby="james woolfenden"
  # }
}
