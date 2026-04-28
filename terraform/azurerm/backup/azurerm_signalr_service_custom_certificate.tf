resource "azurerm_signalr_service_custom_certificate" "pike_gen" {
  name                  = "example-cert"
  signalr_service_id    = "pike"
  custom_certificate_id = "pike"

  depends_on = ["pike"]
}
