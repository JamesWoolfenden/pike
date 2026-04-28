resource "azurerm_network_manager" "pike_gen" {
  name                = "example-network-manager"
  location            = "pike"
  resource_group_name = "pike"
  scope {
    subscription_ids = ["pike"]
  }
  scope_accesses = ["Connectivity", "SecurityAdmin"]
  description    = "example network manager"
  tags = {
    foo = "bar"
  }
}
