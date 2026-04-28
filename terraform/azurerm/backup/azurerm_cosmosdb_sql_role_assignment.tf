resource "azurerm_cosmosdb_sql_role_assignment" "pike_gen" {
  name                = "736180af-7fbc-4c7f-9004-22735173c1c3"
  resource_group_name = "pike"
  account_name        = "pike"
  role_definition_id  = "pike"
  principal_id        = "pike"
  scope               = "pike"
}
