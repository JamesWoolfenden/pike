resource "azurerm_subnet_route_table_association" "pike_gen" {
  subnet_id      = "pike"
  route_table_id = "pike"
}
