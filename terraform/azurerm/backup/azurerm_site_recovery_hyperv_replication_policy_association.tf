resource "azurerm_site_recovery_hyperv_replication_policy_association" "pike_gen" {
  name           = "example-association"
  hyperv_site_id = "pike"
  policy_id      = "pike"
}
