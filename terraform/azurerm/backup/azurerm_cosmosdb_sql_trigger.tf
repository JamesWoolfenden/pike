resource "azurerm_cosmosdb_sql_trigger" "pike_gen" {
  name         = "test-trigger"
  container_id = "pike"
  body         = "function trigger(){}"
  operation    = "Delete"
  type         = "Post"
}
