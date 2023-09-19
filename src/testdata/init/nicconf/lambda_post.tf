module "lambda_post" {
  source = "git::https://github.com/terraform-aws-modules/terraform-aws-lambda.git?ref=2e9aaa2d37d61299bfdaa8b919a75cb37f4726b7"

  function_name = "${random_pet.this.id}-lambda-post"
  description   = "My awesome Python lambda function"
  handler       = "index.lambda_handler"
  runtime       = "python3.8"
  publish       = true

  create_package = false
  s3_existing_package = {
    bucket = "fixtures"
    key    = "python3.8-zip/existing_package.zip"
  }

  # Free TACOS don't have Python available, so we can't build natively there.
  #  source_path = "../src/python-function"
  #  hash_extra  = "post"

  attach_tracing_policy    = true
  attach_policy_statements = true

  policy_statements = {
    dynamodb_write = {
      effect    = "Allow",
      actions   = ["dynamodb:PutItem"],
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
