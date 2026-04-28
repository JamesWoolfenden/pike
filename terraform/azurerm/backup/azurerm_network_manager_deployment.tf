resource "azurerm_network_manager_deployment" "pike_gen" {
  network_manager_id = "pike"
  location           = "eastus"
  scope_access       = "Connectivity"
  configuration_ids  = ["pike"]
}
