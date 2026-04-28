resource "azurerm_network_manager_security_admin_configuration" "pike_gen" {
  name                                          = "example-admin-conf"
  network_manager_id                            = "pike"
  description                                   = "example admin conf"
  apply_on_network_intent_policy_based_services = ["None"]
}
