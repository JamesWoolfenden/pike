resource "azurerm_machine_learning_workspace" "pike_gen" {
  name                    = "example-workspace"
  location                = "pike"
  resource_group_name     = "pike"
  application_insights_id = "pike"
  key_vault_id            = "pike"
  storage_account_id      = "pike"

  identity {
    type = "SystemAssigned"
  }
}
