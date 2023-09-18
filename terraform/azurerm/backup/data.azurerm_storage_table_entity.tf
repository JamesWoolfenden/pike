data "azurerm_storage_table_entity" "pike" {
  table_name           = "pike"
  storage_account_name = "pike"
  partition_key        = "example-partition-key"
  row_key              = "example-row-key"
}
