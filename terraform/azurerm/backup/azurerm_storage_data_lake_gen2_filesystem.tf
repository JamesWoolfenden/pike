resource "azurerm_storage_data_lake_gen2_filesystem" "pike_gen" {
  name               = "example"
  storage_account_id = "pike"

  properties = {
    hello = "aGVsbG8="
  }
}
