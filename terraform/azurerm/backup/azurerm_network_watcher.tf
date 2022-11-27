resource "azurerm_network_watcher" "pike" {
  name                = "production-nwwatcher"
  location            = "uksouth"
  resource_group_name = "pike"
  tags = {
    pike = "permission"
  }
}
