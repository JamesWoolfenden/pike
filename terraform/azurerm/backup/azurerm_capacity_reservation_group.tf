resource "azurerm_capacity_reservation_group" "pike_gen" {
  name                = "example-capacity-reservation-group"
  resource_group_name = "pike"
  location            = "pike"
}
