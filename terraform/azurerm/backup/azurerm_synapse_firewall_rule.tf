resource "azurerm_synapse_firewall_rule" "pike_gen" {
  name                 = "AllowAll"
  synapse_workspace_id = "pike"
  start_ip_address     = "0.0.0.0"
  end_ip_address       = "255.255.255.255"
}
