resource "azurerm_iothub_consumer_group" "pike_gen" {
  name                   = "terraform"
  iothub_name            = "pike"
  eventhub_endpoint_name = "events"
  resource_group_name    = "pike"
}
