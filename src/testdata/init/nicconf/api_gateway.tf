module "api_gateway" {
  source = "git::https://github.com/terraform-aws-modules/terraform-aws-apigateway-v2.git?ref=1e4efc48ab3dfe72d79e80a9c27e4b1bb096e362"

  name          = "${random_pet.this.id}-http"
  description   = "My awesome HTTP API Gateway"
  protocol_type = "HTTP"

  create_api_domain_name = false

  integrations = {
    "GET /" = {
      lambda_arn             = module.lambda_get.lambda_function_arn
      payload_format_version = "2.0"
    }

    "POST /" = {
      lambda_arn             = module.lambda_post.lambda_function_arn
      payload_format_version = "2.0"
    }

    "$default" = {
      lambda_arn = module.lambda_get.lambda_function_arn
    }
  }
}
