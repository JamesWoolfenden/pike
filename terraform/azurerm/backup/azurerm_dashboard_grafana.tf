resource "azurerm_dashboard_grafana" "pike_gen" {
  name                              = "example-dg"
  resource_group_name               = "pike"
  location                          = "West Europe"
  grafana_major_version             = 12
  api_key_enabled                   = true
  deterministic_outbound_ip_enabled = true
  public_network_access_enabled     = false
  sku                               = "Standard"
  sku_size                          = "X1"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    key = "value"
  }
}
