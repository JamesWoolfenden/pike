resource "azurerm_subnet_nat_gateway_association" "pike_gen" {
  subnet_id      = "pike"
  nat_gateway_id = "pike"
}
