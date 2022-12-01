resource "azurerm_web_pubsub" "pike" {
  name                = "tfex-webpubsub"
  location            = "uksouth"
  resource_group_name = "pike"

  sku      = "Standard_S1"
  capacity = 1

  public_network_access_enabled = false

  live_trace {
    enabled                   = true
    messaging_logs_enabled    = true
    connectivity_logs_enabled = false
  }

  identity {
    type = "SystemAssigned"
  }
}
