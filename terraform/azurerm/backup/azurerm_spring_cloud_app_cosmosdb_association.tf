resource "azurerm_spring_cloud_app_cosmosdb_association" "pike_gen" {
  name                = "example-bind"
  spring_cloud_app_id = "pike"
  cosmosdb_account_id = "pike"
  api_type            = "table"
}
