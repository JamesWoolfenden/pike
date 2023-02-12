data "aws_service_discovery_dns_namespace" "pike" {
  name = "example.terraform.local"
  type = "DNS_PRIVATE"
}
