resource "azurerm_cdn_profile" "pike_gen" {
  name                = "exampleCdnProfile"
  location            = "pike"
  resource_group_name = "pike"
  sku                 = "Standard_Microsoft"

  tags = {
    environment = "Production"
    cost_center = "MSFT"
  }
}
