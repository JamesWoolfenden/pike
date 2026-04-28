data "azurerm_netapp_snapshot_policy" "pike_gen" {
  resource_group_name = "acctestRG"
  account_name        = "acctestnetappaccount"
  name                = "example-snapshot-policy"
}
