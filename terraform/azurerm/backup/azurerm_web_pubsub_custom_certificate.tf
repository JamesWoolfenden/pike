resource "azurerm_web_pubsub_custom_certificate" "pike_gen" {
  name                  = "example-cert"
  web_pubsub_id         = "pike"
  custom_certificate_id = "pike"

  depends_on = ["pike"]
}
