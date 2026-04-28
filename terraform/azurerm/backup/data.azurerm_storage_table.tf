data "azurerm_storage_table" "pike_gen" {
  name                 = "example-table-name"
  storage_account_name = "example-storage-account-name"
}
