resource "azurerm_confidential_ledger" "pike_gen" {
  name                = "example-ledger"
  resource_group_name = "pike"
  location            = "pike"
  ledger_type         = "Private"

  azuread_based_service_principal {
    principal_id     = "pike"
    tenant_id        = "pike"
    ledger_role_name = "Administrator"
  }
}
