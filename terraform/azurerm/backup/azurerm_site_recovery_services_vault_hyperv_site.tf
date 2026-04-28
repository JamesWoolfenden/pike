resource "azurerm_site_recovery_services_vault_hyperv_site" "pike_gen" {
  name              = "example-site"
  recovery_vault_id = "pike"
}
