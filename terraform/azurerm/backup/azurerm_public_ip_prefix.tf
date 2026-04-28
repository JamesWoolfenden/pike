resource "azurerm_public_ip_prefix" "pike_gen" {
  name                = "acceptanceTestPublicIpPrefix1"
  location            = "pike"
  resource_group_name = "pike"

  prefix_length = 31

  tags = {
    environment = "Production"
  }
}
