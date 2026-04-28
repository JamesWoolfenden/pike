resource "azurerm_api_management_identity_provider_google" "pike_gen" {
  resource_group_name = "pike"
  api_management_name = "pike"
  client_id           = "00000000.apps.googleusercontent.com"
}
