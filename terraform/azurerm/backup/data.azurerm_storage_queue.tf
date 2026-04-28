data "azurerm_storage_queue" "pike_gen" {
  name                 = "example-queue-name"
  storage_account_name = "example-storage-account-name"
}
