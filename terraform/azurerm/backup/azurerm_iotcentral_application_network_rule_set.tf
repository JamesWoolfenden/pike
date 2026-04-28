resource "azurerm_iotcentral_application_network_rule_set" "pike_gen" {
  iotcentral_application_id = "pike"

  ip_rule {
    name    = "rule1"
    ip_mask = "10.0.1.0/24"
  }

  ip_rule {
    name    = "rule2"
    ip_mask = "10.1.1.0/24"
  }
}
