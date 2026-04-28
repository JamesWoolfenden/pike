resource "azurerm_api_management_standalone_gateway" "pike_gen" {
  name                 = "example-gateway-flexible"
  resource_group_name  = "pike"
  location             = "pike"
  virtual_network_type = "External"
  backend_subnet_id    = "pike"

  sku {
    capacity = 1
    name     = "WorkspaceGatewayPremium"
  }

  tags = {
    Hello = "World"
  }
}
