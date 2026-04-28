resource "azurerm_network_manager_connectivity_configuration" "pike_gen" {
  name                  = "example-connectivity-conf"
  network_manager_id    = "pike"
  connectivity_topology = "HubAndSpoke"
  applies_to_group {
    group_connectivity = "DirectlyConnected"
    network_group_id   = "pike"
  }

  applies_to_group {
    group_connectivity = "DirectlyConnected"
    network_group_id   = "pike"
  }

  hub {
    resource_id   = "pike"
    resource_type = "Microsoft.Network/virtualNetworks"
  }
}
