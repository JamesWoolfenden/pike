resource "azurerm_cosmosdb_sql_dedicated_gateway" "pike_gen" {
  cosmosdb_account_id = "pike"
  instance_count      = 1
  instance_size       = "Cosmos.D4s"
}
