resource "azurerm_maps_account" "pike_gen" {
  name                         = "example-maps-account"
  resource_group_name          = "pike"
  sku_name                     = "S1"
  local_authentication_enabled = true

  tags = {
    environment = "Test"
  }
}
