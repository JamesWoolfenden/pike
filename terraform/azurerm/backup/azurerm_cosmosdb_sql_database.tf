resource "azurerm_cosmosdb_sql_database" "pike_gen" {
  name                = "tfex-cosmos-sql-db"
  resource_group_name = "pike"
  account_name        = "pike"
  throughput          = 400
}
