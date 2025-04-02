data "aws_apigateway_api_keys" "pike" {
}

output "aws_apigateway_api_keys" {
  value = data.aws_apigateway_api_keys.pike
}
