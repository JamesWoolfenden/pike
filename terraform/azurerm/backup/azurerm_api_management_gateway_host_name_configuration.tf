resource "azurerm_api_management_gateway_host_name_configuration" "pike_gen" {
  name              = "example-host-name-configuration"
  api_management_id = "pike"
  gateway_name      = "pike"

  certificate_id                     = "pike"
  host_name                          = "example-host-name"
  request_client_certificate_enabled = true
  http2_enabled                      = true
  tls10_enabled                      = true
  tls11_enabled                      = false
}
