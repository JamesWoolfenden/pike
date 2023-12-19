data "azurerm_site_recovery_fabric" "pike" {
  resource_group_name = "pike"
  recovery_vault_name = "pike"
  name                = "pike"
}
