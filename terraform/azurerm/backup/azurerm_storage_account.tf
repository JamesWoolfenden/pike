resource "azurerm_storage_account" "pike" {
  name                     = "jgw08202722"
  location                 = "uksouth"
  resource_group_name      = "NetworkWatcherRG"
  account_tier             = "Standard"
  account_replication_type = "GRS"

  tags = {
    pike = "permissions"
  }
}
