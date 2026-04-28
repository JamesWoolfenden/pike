data "azurerm_cosmosdb_sql_database" "pike_gen" {
  name                = "tfex-cosmosdb-sql-database"
  resource_group_name = "tfex-cosmosdb-sql-database-rg"
  account_name        = "tfex-cosmosdb-sql-database-account-name"
}
