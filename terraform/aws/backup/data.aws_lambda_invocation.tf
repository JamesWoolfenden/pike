data "aws_lambda_invocation" "pike" {
  function_name = "helloworld"
  input         = "[]"
}


output "result" {
  value=data.aws_lambda_invocation.pike
}