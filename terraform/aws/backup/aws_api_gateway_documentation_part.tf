resource "aws_api_gateway_documentation_part" "pike" {
  location {
    type   = "METHOD"
    method = "GET"
    path   = "/example"
  }

  properties  = "{\"description\":\"Example description\"}"
  rest_api_id = "7n300zvss6"
}
