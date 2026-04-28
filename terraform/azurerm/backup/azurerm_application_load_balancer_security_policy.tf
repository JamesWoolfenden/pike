resource "azurerm_application_load_balancer_security_policy" "pike_gen" {
  name                               = "example-albsp"
  application_load_balancer_id       = "pike"
  location                           = "pike"
  web_application_firewall_policy_id = "pike"
}
