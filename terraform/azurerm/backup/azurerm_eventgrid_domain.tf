resource "azurerm_eventgrid_domain" "pike_gen" {
  name                = "my-eventgrid-domain"
  location            = "pike"
  resource_group_name = "pike"

  tags = {
    environment = "Production"
  }
}
