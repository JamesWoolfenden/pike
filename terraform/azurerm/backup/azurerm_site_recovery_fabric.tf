resource "azurerm_site_recovery_fabric" "pike_gen" {
  name                = "primary-fabric"
  resource_group_name = "pike"
  recovery_vault_name = "pike"
  location            = "pike"
}
