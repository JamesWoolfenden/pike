resource "azurerm_synapse_sql_pool" "pike_gen" {
  name                 = "examplesqlpool"
  synapse_workspace_id = "pike"
  sku_name             = "DW100c"
  create_mode          = "Default"
  storage_account_type = "GRS"
}
