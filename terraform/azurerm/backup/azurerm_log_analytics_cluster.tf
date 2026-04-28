resource "azurerm_log_analytics_cluster" "pike_gen" {
  name                = "example-cluster"
  resource_group_name = "pike"
  location            = "pike"

  identity {
    type = "SystemAssigned"
  }
}
