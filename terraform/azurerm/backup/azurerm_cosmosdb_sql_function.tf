resource "azurerm_cosmosdb_sql_function" "pike_gen" {
  name         = "test-function"
  container_id = "pike"
  body         = "function trigger(){}"
}
