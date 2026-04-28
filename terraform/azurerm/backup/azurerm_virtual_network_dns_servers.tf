resource "azurerm_virtual_network_dns_servers" "pike_gen" {
  virtual_network_id = "pike"
  dns_servers        = ["10.7.7.2", "10.7.7.7", "10.7.7.1"]
}
