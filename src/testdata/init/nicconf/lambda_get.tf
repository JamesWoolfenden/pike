module "lambda_get" {
  source = "git::https://github.com/terraform-aws-modules/terraform-aws-lambda.git?ref=2e9aaa2d37d61299bfdaa8b919a75cb37f4726b7"

  function_name = "${random_pet.this.id}-lambda-get"
  description   = "My awesome Python lambda function"
  handler       = "index.lambda_handler"
  runtime       = "python3.8"
  publish       = true

  create_package = false
  s3_existing_package = {
    bucket = "fixtures"
    key    = "python3.8-zip/existing_package.zip"
  }

  attach_tracing_policy    = true
  attach_policy_statements = true

  policy_statements = {
    dynamodb_read = {
      effect    = "Allow",
      actions   = ["dynamodb:GetItem"],
      resources = [module.dynamodb_table.dynamodb_table_arn]
    }
  }

  allowed_triggers = {
    AllowExecutionFromAPIGateway = {
      service    = "apigateway"
      source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*/*/*"
    }
  }
}
