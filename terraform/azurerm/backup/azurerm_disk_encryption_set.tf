resource "azurerm_disk_encryption_set" "pike" {
  name                = "pike"
  resource_group_name = "pike"
  location            = "uksouth"
  key_vault_key_id    = data.azurerm_key_vault_key.pike.id

  identity {
    type = "SystemAssigned"
  }
}

data "azurerm_key_vault_key" "pike" {

  key_vault_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.KeyVault/vaults/pike"
  name         = "pike"
}
