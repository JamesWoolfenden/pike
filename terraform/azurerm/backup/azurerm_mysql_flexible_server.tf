resource "azurerm_mysql_flexible_server" "pike_gen" {
  name                  = "example-fs"
  resource_group_name   = "pike"
  location              = "pike"
  administrator_login   = "psqladmin"
  backup_retention_days = 7
  delegated_subnet_id   = "pike"
  private_dns_zone_id   = "pike"
  sku_name              = "GP_Standard_D2ds_v4"

  depends_on = ["pike"]
}
