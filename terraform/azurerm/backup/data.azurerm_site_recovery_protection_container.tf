data "azurerm_site_recovery_protection_container" "pike" {
  resource_group_name  = "pike"
  recovery_fabric_name = "pike"
  recovery_vault_name  = "pike"
  name                 = "pike"
}
