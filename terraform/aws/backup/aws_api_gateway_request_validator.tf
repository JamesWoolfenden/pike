resource "aws_api_gateway_request_validator" "pike" {
  name                        = "example"
  rest_api_id                 = "7n300zvss6"
  validate_request_body       = true
  validate_request_parameters = true
}
