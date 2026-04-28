resource "azurerm_graph_services_account" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  application_id      = "pike"
  tags = {
    environment = "Production"
  }
}
