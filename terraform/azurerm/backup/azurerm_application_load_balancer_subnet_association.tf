resource "azurerm_application_load_balancer_subnet_association" "pike_gen" {
  name                         = "example"
  application_load_balancer_id = "pike"
  subnet_id                    = "pike"
}
