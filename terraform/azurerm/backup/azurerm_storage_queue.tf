resource "azurerm_storage_queue" "pike_gen" {
  name                 = "mysamplequeue"
  storage_account_name = "pike"
}
