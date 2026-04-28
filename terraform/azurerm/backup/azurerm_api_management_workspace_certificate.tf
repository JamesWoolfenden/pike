resource "azurerm_api_management_workspace_certificate" "pike_gen" {
  name                        = "example-cert"
  api_management_workspace_id = "pike"
  certificate_data_base64     = filebase64("example.pfx")
}
