resource "azurerm_mssql_database_extended_auditing_policy" "pike_gen" {
  database_id       = "pike"
  storage_endpoint  = "pike"
  retention_in_days = 6
}
