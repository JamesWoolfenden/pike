resource "aws_api_gateway_api_key" "MyDemoApiKey" {
  name = "demo"
  tags = {
    Pike    = "permissions"
    nothing = "new"
  }
}
