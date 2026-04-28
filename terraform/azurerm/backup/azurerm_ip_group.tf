resource "azurerm_ip_group" "pike_gen" {
  name                = "example1-ipgroup"
  location            = "pike"
  resource_group_name = "pike"

  cidrs = ["192.168.0.1", "172.16.240.0/20", "10.48.0.0/12"]

  tags = {
    environment = "Production"
  }
}
