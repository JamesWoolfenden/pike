resource "azurerm_log_analytics_workspace" "pike" {
  name                = "pike"
  location            = "uksouth"
  resource_group_name = "NetworkWatcherRG"
  sku                 = "PerGB2018"
  retention_in_days   = 30
}
