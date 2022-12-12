resource "azurerm_search_service" "pike" {
  location                      = "uksouth"
  name                          = "pike"
  resource_group_name           = "pike"
  sku                           = "basic"
  public_network_access_enabled = false
}
