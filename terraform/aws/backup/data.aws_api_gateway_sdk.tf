data "aws_api_gateway_sdk" "pike" {
  rest_api_id = "7n300zvss6"
  stage_name  = "prod"
  sdk_type    = "android"
}
