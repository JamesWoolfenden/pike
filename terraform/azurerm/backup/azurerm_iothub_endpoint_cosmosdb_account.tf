resource "azurerm_iothub_endpoint_cosmosdb_account" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  iothub_id           = "pike"
  container_name      = "pike"
  database_name       = "pike"
  endpoint_uri        = "pike"
  primary_key         = "pike"
  secondary_key       = "pike"
}
