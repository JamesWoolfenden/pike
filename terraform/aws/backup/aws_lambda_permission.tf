resource "aws_lambda_alias" "test_alias" {
  name             = "testalias"
  description      = "a simple description"
  function_name    = "updatepolicy"
  function_version = "$LATEST"
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "updatepolicy"
  principal     = "events.amazonaws.com"
  source_arn    = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily"
  qualifier     = aws_lambda_alias.test_alias.name
}
