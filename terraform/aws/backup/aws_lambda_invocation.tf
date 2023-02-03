resource "aws_lambda_invocation" "pike" {
  function_name = "helloworld"

  input = jsonencode({
    key1 = "value1"
    key2 = "value2"
  })
}
