resource "azurerm_mssql_server_extended_auditing_policy" "pike_gen" {
  server_id         = "pike"
  storage_endpoint  = "pike"
  retention_in_days = 6
}
