resource "aws_api_gateway_rest_api" "pikeexample" {
  #  body = jsonencode({
  #    openapi = "3.0.1"
  #    info = {
  #      title   = "examples"
  #      version = "1.0"
  #    }
  #    paths = {
  #      "/path1" = {
  #        get = {
  #          x-amazon-apigateway-integration = {
  #            httpMethod           = "GET"
  #            payloadFormatVersion = "1.0"
  #            type                 = "HTTP_PROXY"
  #            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
  #          }
  #        }
  #      }
  #    }
  #  })

  name = "example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}
