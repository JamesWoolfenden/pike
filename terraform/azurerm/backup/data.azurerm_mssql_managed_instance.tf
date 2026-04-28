data "azurerm_mssql_managed_instance" "pike_gen" {
  name                = "managedsqlinstance"
  resource_group_name = "pike"
}
