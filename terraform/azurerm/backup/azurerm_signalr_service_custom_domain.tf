resource "azurerm_signalr_service_custom_domain" "pike_gen" {
  name                          = "example-domain"
  signalr_service_id            = "pike"
  domain_name                   = "tftest.com"
  signalr_custom_certificate_id = "pike"
}
