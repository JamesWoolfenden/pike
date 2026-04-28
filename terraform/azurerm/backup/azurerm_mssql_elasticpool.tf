resource "azurerm_mssql_elasticpool" "pike_gen" {
  name                = "test-epool"
  resource_group_name = "pike"
  location            = "pike"
  server_name         = "pike"
  license_type        = "LicenseIncluded"
  max_size_gb         = 756

  sku {
    name     = "BasicPool"
    tier     = "Basic"
    family   = "Gen4"
    capacity = 4
  }

  per_database_settings {
    min_capacity = 0.25
    max_capacity = 4
  }
}
