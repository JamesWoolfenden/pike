resource "azurerm_api_management_openid_connect_provider" "pike_gen" {
  name                = "example-provider"
  api_management_name = "pike"
  resource_group_name = "pike"
  client_id           = "00001111-2222-3333-4444-555566667777"
  display_name        = "Example Provider"
  metadata_endpoint   = "https://example.com/example"
}
