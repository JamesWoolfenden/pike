resource "azurerm_public_ip" "pike_gen" {
  name                = "acceptanceTestPublicIp1"
  resource_group_name = "pike"
  location            = "pike"
  allocation_method   = "Static"

  tags = {
    environment = "Production"
  }
}
