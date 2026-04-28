resource "azurerm_data_factory_linked_service_odata" "pike_gen" {
  name            = "anonymous"
  data_factory_id = "pike"
  url             = "https://services.odata.org/v4/TripPinServiceRW/People"
}
