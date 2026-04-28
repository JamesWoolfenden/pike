resource "azurerm_iothub_endpoint_eventhub" "pike_gen" {
  resource_group_name = "pike"
  iothub_id           = "pike"
  name                = "example"

}
