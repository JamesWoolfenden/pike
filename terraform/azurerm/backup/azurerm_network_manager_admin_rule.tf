resource "azurerm_network_manager_admin_rule" "pike_gen" {
  name                     = "example-admin-rule"
  admin_rule_collection_id = "pike"
  action                   = "Deny"
  direction                = "Outbound"
  priority                 = 1
  protocol                 = "Tcp"
  source_port_ranges       = ["80", "1024-65535"]
  destination_port_ranges  = ["80"]
  source {
    address_prefix_type = "ServiceTag"
    address_prefix      = "Internet"
  }
  destination {
    address_prefix_type = "IPPrefix"
    address_prefix      = "10.1.0.1"
  }
  destination {
    address_prefix_type = "IPPrefix"
    address_prefix      = "10.0.0.0/24"
  }
  description = "example admin rule"
}
