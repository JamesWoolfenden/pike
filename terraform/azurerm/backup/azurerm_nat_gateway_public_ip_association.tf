resource "azurerm_nat_gateway_public_ip_association" "pike_gen" {
  nat_gateway_id       = "pike"
  public_ip_address_id = "pike"
}
