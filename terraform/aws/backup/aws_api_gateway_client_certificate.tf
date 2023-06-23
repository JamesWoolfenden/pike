resource "aws_api_gateway_client_certificate" "pike" {
  description = "pike"
  tags = {
    pike = "permissions"
  }
}
