resource "azurerm_netapp_snapshot" "pike_gen" {
  name                = "example-netappsnapshot"
  account_name        = "pike"
  pool_name           = "pike"
  volume_name         = "pike"
  location            = "pike"
  resource_group_name = "pike"
}
