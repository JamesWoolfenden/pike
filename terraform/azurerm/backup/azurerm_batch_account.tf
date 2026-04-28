resource "azurerm_batch_account" "pike_gen" {
  name                                = "testbatchaccount"
  resource_group_name                 = "pike"
  location                            = "pike"
  pool_allocation_mode                = "BatchService"
  storage_account_id                  = "pike"
  storage_account_authentication_mode = "StorageKeys"

  tags = {
    env = "test"
  }
}
