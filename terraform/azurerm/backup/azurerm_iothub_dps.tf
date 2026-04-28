resource "azurerm_iothub_dps" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  allocation_policy   = "Hashed"

  sku {
    name     = "S1"
    capacity = "1"
  }
}
