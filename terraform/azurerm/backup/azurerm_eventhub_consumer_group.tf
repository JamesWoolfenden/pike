resource "azurerm_eventhub_consumer_group" "pike_gen" {
  name                = "acceptanceTestEventHubConsumerGroup"
  namespace_name      = "pike"
  eventhub_name       = "pike"
  resource_group_name = "pike"
  user_metadata       = "some-meta-data"
}
