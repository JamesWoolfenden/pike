resource "azurerm_eventgrid_topic" "pike_gen" {
  name                = "my-eventgrid-topic"
  location            = "pike"
  resource_group_name = "pike"

  tags = {
    environment = "Production"
  }
}
