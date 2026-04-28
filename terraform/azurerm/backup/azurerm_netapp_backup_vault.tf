resource "azurerm_netapp_backup_vault" "pike_gen" {
  name                = "example-netappbackupvault"
  resource_group_name = "pike"
  location            = "pike"
  account_name        = "pike"
}
