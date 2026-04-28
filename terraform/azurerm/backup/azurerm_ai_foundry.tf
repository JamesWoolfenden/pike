resource "azurerm_ai_foundry" "pike_gen" {
  name                = "exampleaihub"
  location            = "pike"
  resource_group_name = "pike"
  storage_account_id  = "pike"
  key_vault_id        = "pike"

  identity {
    type = "SystemAssigned"
  }
}
