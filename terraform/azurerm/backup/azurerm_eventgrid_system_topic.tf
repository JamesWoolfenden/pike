resource "azurerm_eventgrid_system_topic" "pike_gen" {
  name                = "example-topic"
  resource_group_name = "pike"
  location            = "pike"
  source_resource_id  = "pike"
  topic_type          = "Microsoft.Storage.StorageAccounts"
}
