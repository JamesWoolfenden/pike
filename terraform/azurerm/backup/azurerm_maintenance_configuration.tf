resource "azurerm_maintenance_configuration" "pike_gen" {
  name                = "example-mc"
  resource_group_name = "pike"
  location            = "pike"
  scope               = "SQLDB"

  tags = {
    Env = "prod"
  }
}
