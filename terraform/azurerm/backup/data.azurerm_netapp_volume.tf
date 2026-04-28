data "azurerm_netapp_volume" "pike_gen" {
  resource_group_name = "acctestRG"
  account_name        = "acctestnetappaccount"
  pool_name           = "acctestnetapppool"
  name                = "example-volume"
}
