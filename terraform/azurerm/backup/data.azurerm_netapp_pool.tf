data "azurerm_netapp_pool" "pike_gen" {
  resource_group_name = "acctestRG"
  account_name        = "acctestnetappaccount"
  name                = "acctestnetapppool"
}
