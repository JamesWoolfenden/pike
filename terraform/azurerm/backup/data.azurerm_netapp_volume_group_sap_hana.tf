data "azurerm_netapp_volume_group_sap_hana" "pike_gen" {
  name                = "existing application volume group name"
  resource_group_name = "resource group name where the account and volume group belong to"
  account_name        = "existing account where the application volume group belong to"
}
