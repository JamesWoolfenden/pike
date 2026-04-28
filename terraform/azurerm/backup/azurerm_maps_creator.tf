resource "azurerm_maps_creator" "pike_gen" {
  name            = "example-maps-creator"
  maps_account_id = "pike"
  location        = "pike"
  storage_units   = 1

  tags = {
    environment = "Test"
  }
}
