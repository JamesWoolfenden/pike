resource "azurerm_web_pubsub_custom_domain" "pike_gen" {
  name                             = "example-domain"
  domain_name                      = "tftest.com"
  web_pubsub_id                    = "pike"
  web_pubsub_custom_certificate_id = "pike"
}
