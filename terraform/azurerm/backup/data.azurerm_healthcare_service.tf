data "azurerm_healthcare_service" "pike_gen" {
  name                = "example-healthcare_service"
  resource_group_name = "example-resources"
  location            = "westus2"
}
