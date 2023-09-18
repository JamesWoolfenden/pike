data "azurerm_storage_management_policy" "pike" {
  storage_account_id = data.azurerm_storage_account.pike.id
}
