resource "aws_service_discovery_public_dns_namespace" "pike" {
  name        = "services.pike.com"
  description = "example"
}
