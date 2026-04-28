resource "azurerm_data_factory_linked_service_azure_search" "pike_gen" {
  name               = "example"
  data_factory_id    = "pike"
  url                = join("", ["https://", "pike", ".search.windows.net"])
  search_service_key = "pike"
}
