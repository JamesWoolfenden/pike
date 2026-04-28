resource "azurerm_mssql_managed_instance" "pike_gen" {
  name                = "managedsqlinstance"
  resource_group_name = "pike"
  location            = "pike"

  license_type       = "BasePrice"
  sku_name           = "GP_Gen5"
  storage_size_in_gb = 32
  subnet_id          = "pike"
  vcores             = 4

  administrator_login = "mradministrator"

  depends_on = [
    azurerm_subnet_network_security_group_association.example,
    azurerm_subnet_route_table_association.example,
  ]
}
