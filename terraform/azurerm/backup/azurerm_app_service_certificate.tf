resource "azurerm_app_service_certificate" "pike_gen" {
  name                = "example-cert"
  resource_group_name = "pike"
  location            = "pike"
  pfx_blob            = filebase64("certificate.pfx")
}
