resource "azurerm_synapse_sql_pool_extended_auditing_policy" "pike_gen" {
  sql_pool_id       = "pike"
  storage_endpoint  = "pike"
  retention_in_days = 6
}
