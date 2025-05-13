data "aws_lambda_function_url" "pike" {
  function_name = "pike"
}

output "aws_lambda_function_url" {
  value = data.aws_lambda_function_url.pike
}
