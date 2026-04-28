resource "azurerm_site_recovery_vmware_replication_policy_association" "pike_gen" {
  name              = "example-association"
  recovery_vault_id = "pike"
  policy_id         = "pike"
}
