resource "azurerm_capacity_reservation" "pike_gen" {
  name                          = "example-capacity-reservation"
  capacity_reservation_group_id = "pike"
  sku {
    name     = "Standard_D2s_v3"
    capacity = 1
  }
}
