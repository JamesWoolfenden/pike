resource "azurerm_nat_gateway" "pike_gen" {
  name                    = "nat-gateway"
  location                = "pike"
  resource_group_name     = "pike"
  sku_name                = "Standard"
  idle_timeout_in_minutes = 10
  zones                   = ["1"]
}
