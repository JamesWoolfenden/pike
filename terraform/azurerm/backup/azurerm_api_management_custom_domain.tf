resource "azurerm_api_management_custom_domain" "pike_gen" {
  api_management_id = "pike"

  gateway {
    host_name                = "api.example.com"
    key_vault_certificate_id = "pike"
  }

  developer_portal {
    host_name                = "portal.example.com"
    key_vault_certificate_id = "pike"
  }
}
