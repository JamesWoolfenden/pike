resource "azurerm_storage_account_customer_managed_key" "pike" {
  key_name           = "pike"
  key_vault_id       = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.KeyVault/vaults/pike"
  storage_account_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Storage/storageAccounts/pike"
}
