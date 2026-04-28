data "azurerm_cosmosdb_sql_role_definition" "pike_gen" {
  resource_group_name = "tfex-cosmosdb-sql-role-definition-rg"
  account_name        = "tfex-cosmosdb-sql-role-definition-account-name"
  role_definition_id  = "00000000-0000-0000-0000-000000000000"
}
