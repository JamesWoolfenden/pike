resource "azurerm_custom_provider" "pike_gen" {
  name                = "example_provider"
  location            = "pike"
  resource_group_name = "pike"

  resource_type {
    name     = "dEf1"
    endpoint = "https://testendpoint.com/"
  }
}
