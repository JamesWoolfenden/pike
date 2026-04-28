resource "azurerm_databricks_virtual_network_peering" "pike_gen" {
  name                = "databricks-vnet-peer"
  resource_group_name = "pike"
  workspace_id        = "pike"

  remote_address_space_prefixes = "pike"
  remote_virtual_network_id     = "pike"
  allow_virtual_network_access  = true
}
