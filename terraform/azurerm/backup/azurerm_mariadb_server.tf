resource "azurerm_mariadb_server" "pike" {
  name                = "mariadb-svr-jgw"
  location            = "uksouth"
  resource_group_name = "pike"

  sku_name = "B_Gen5_2"

  storage_mb                   = 51200
  backup_retention_days        = 7
  geo_redundant_backup_enabled = false

  administrator_login          = "acctestun"
  administrator_login_password = "H@Sh1CoR3!"
  version                      = "10.2"
  ssl_enforcement_enabled      = true
  tags = {
    pike = "permissions"
  }
}
