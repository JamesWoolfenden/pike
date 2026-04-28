resource "azurerm_iothub_endpoint_servicebus_topic" "pike_gen" {
  resource_group_name = "pike"
  iothub_id           = "pike"
  name                = "example"

}
