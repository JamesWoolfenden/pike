data "azurerm_eventhub_consumer_group" "pike_gen" {
  name                = "pike"
  namespace_name      = "pike"
  eventhub_name       = "pike"
  resource_group_name = "pike"
}
