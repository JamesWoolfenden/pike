data "aws_api_gateway_export" "pike" {
  export_type = "swagger"
  rest_api_id = "7n300zvss6"
  stage_name  = "prod"
}
