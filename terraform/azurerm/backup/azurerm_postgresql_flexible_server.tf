resource "azurerm_postgresql_flexible_server" "pike_gen" {
  name                          = "example-psqlflexibleserver"
  resource_group_name           = "pike"
  location                      = "pike"
  version                       = "12"
  delegated_subnet_id           = "pike"
  private_dns_zone_id           = "pike"
  public_network_access_enabled = false
  administrator_login           = "psqladmin"
  zone                          = "1"

  storage_mb   = 32768
  storage_tier = "P4"

  sku_name   = "B_Standard_B1ms"
  depends_on = ["pike"]

}
