resource "azurerm_eventgrid_namespace" "pike_gen" {
  name                = "my-eventgrid-namespace"
  location            = "pike"
  resource_group_name = "pike"

  tags = {
    environment = "Production"
  }
}
