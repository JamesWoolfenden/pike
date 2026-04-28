resource "azurerm_netapp_backup_policy" "pike_gen" {
  name                = "example-netappbackuppolicy"
  resource_group_name = "pike"
  location            = "pike"
  account_name        = "pike"
  enabled             = true
}
