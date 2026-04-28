resource "azurerm_kusto_cosmosdb_data_connection" "pike_gen" {
  name                  = "examplekcdcd"
  location              = "pike"
  cosmosdb_container_id = "pike"
  kusto_database_id     = "pike"
  managed_identity_id   = "pike"
  table_name            = "TestTable"
  mapping_rule_name     = "TestMapping"
  retrieval_start_date  = "2023-06-26T12:00:00.6554616Z"
}
