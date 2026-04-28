resource "azurerm_security_center_automation" "pike_gen" {
  name                = "example-automation"
  location            = "pike"
  resource_group_name = "pike"

  action {
    type        = "EventHub"
    resource_id = "pike"
  }

  source {
    event_source = "Alerts"
    rule_set {
      rule {
        property_path  = "properties.metadata.severity"
        operator       = "Equals"
        expected_value = "High"
        property_type  = "String"
      }
    }
  }

  scopes = ["/subscriptions/${data.azurerm_client_config.current.subscription_id}"]
}
