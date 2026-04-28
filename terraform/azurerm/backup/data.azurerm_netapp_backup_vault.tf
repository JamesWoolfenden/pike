data "azurerm_netapp_backup_vault" "pike_gen" {
  resource_group_name = "example-resource-group"
  account_name        = "example-netappaccount"
  name                = "example-backupvault"
}
