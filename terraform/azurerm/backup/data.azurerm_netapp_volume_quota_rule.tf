data "azurerm_netapp_volume_quota_rule" "pike_gen" {
  name      = "exampleQuotaRule"
  volume_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/vol1"
}
