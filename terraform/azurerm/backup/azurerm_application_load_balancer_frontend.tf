resource "azurerm_application_load_balancer_frontend" "pike_gen" {
  name                         = "example"
  application_load_balancer_id = "pike"
}
