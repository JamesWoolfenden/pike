resource "azurerm_cosmosdb_mongo_user_definition" "pike_gen" {
  cosmos_mongo_database_id = "pike"
  username                 = "myUserName"
}
