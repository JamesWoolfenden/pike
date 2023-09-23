data "aws_lambda_alias" "pike" {
  function_name = "message-processor"
  name          = "pike"
}
