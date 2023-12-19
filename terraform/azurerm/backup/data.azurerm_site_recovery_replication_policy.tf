data "azurerm_site_recovery_replication_policy" "pike" {
  resource_group_name = "pike"
  name                = "pike"
  recovery_vault_name = "pike"
}
