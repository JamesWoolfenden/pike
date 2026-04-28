resource "azurerm_data_factory_linked_service_azure_function" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  url             = "https://${data.azurerm_function_app.example.default_hostname}"
  key             = "foo"

}
