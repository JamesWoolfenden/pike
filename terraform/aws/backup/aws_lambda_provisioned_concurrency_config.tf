resource "aws_lambda_provisioned_concurrency_config" "pike" {
  function_name                     = "message-processor"
  provisioned_concurrent_executions = 2
  qualifier                         = "1"
}
