resource "azurerm_communication_service" "pike_gen" {
  name                = "example-communicationservice"
  resource_group_name = "pike"
  data_location       = "United States"
}
