data "aws_apigatewayv2_export" "pike" {
  api_id        = "cz518irapl"
  specification = "OAS30"
  output_type   = "JSON"
}
