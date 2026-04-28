resource "azurerm_netapp_pool" "pike_gen" {
  name                = "example-netapppool"
  account_name        = "pike"
  location            = "pike"
  resource_group_name = "pike"
  service_level       = "Premium"
  size_in_tb          = 4
}
