data "azurerm_netapp_snapshot" "pike_gen" {
  resource_group_name = "acctestRG"
  name                = "acctestnetappsnapshot"
  account_name        = "acctestnetappaccount"
  pool_name           = "acctestnetapppool"
  volume_name         = "acctestnetappvolume"
}
