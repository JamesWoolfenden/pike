data "azurerm_storage_containers" "pike" {
  storage_account_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Storage/storageAccounts/pike"
}

output "container" {
  value = data.azurerm_storage_containers.pike
}
