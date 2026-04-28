resource "azurerm_active_directory_domain_service_trust" "pike_gen" {
  name                   = "example-trust"
  domain_service_id      = "pike"
  trusted_domain_fqdn    = "example.com"
  trusted_domain_dns_ips = ["10.1.0.3", "10.1.0.4"]
}
