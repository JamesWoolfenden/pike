data "aws_lambda_functions" "pike" {

}

output "aws_lambda_functions" {
  value = data.aws_lambda_functions.pike
}
