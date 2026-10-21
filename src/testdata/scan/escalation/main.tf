resource "aws_lambda_function" "pike" {
  function_name = "pike"
  role          = "arn:aws:iam::123456789012:role/pike"
  handler       = "index.handler"
  runtime       = "nodejs18.x"
  filename      = "lambda.zip"
}
