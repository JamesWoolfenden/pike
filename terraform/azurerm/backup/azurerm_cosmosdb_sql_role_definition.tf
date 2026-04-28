resource "azurerm_cosmosdb_sql_role_definition" "pike_gen" {
  role_definition_id  = "84cf3a8b-4122-4448-bce2-fa423cfe0a15"
  resource_group_name = "pike"
  account_name        = "pike"
  name                = "acctestsqlrole"
  assignable_scopes   = ["${azurerm_cosmosdb_account.example.id}/dbs/sales"]

  permissions {
    data_actions = ["Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read"]
  }
}
