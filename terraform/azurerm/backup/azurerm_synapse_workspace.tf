resource "azurerm_synapse_workspace" "pike_gen" {
  name                                 = "example"
  resource_group_name                  = "pike"
  location                             = "pike"
  storage_data_lake_gen2_filesystem_id = "pike"
  sql_administrator_login              = "sqladminuser"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Env = "production"
  }
}
