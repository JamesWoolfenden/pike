resource "aws_wafv2_web_acl_association" "pike" {
  resource_arn = aws_api_gateway_stage.example.arn
  #   resource_arn = aws_cognito_user_pool.pool.arn
  web_acl_arn = aws_wafv2_web_acl.example.arn
}

resource "aws_wafv2_web_acl" "example" {
  name  = "web-acl-association-example"
  scope = "REGIONAL"

  default_action {
    allow {}
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "friendly-metric-name"
    sampled_requests_enabled   = false
  }
}

resource "aws_api_gateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "example"
}

resource "aws_api_gateway_deployment" "example" {
  rest_api_id = aws_api_gateway_rest_api.example.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.example.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "example" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}
# Application Load Balancer
#
# AWS AppSync GraphQL API
#
# Amazon Cognito user pool
# resource "aws_cognito_user_pool" "pool" {
#   name = "mypool"
#   account_recovery_setting {
#     recovery_mechanism {
#       name     = "verified_email"
#       priority = 1
#     }
#
#     recovery_mechanism {
#       name     = "verified_phone_number"
#       priority = 2
#     }
#   }
# }
# AWS App Runner service
#
# AWS Verified Access instance
