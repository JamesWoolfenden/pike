resource "azurerm_netapp_volume_quota_rule" "pike_gen" {
  name              = "example-quota-rule-1"
  location          = "pike"
  volume_id         = "pike"
  quota_target      = "3001"
  quota_size_in_kib = 1024
  quota_type        = "IndividualGroupQuota"
}
