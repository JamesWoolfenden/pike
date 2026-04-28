resource "azurerm_storage_data_lake_gen2_path" "pike_gen" {
  path               = "example"
  filesystem_name    = "pike"
  storage_account_id = "pike"
  resource           = "directory"
}
