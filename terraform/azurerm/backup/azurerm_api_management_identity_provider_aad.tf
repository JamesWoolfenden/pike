resource "azurerm_api_management_identity_provider_aad" "pike_gen" {
  resource_group_name = "pike"
  api_management_name = "pike"
  client_id           = "00000000-0000-0000-0000-000000000000"
  allowed_tenants     = ["00000000-0000-0000-0000-000000000000"]
}
