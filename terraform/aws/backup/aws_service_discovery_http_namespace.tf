resource "aws_service_discovery_http_namespace" "pike" {
  name        = "pike"
  description = "A service discovery http namespace."
  tags = {
    pike = "permissions"
  }
}
