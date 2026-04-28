resource "azurerm_servicebus_namespace_disaster_recovery_config" "pike_gen" {
  name                        = "servicebus-alias-name"
  primary_namespace_id        = "pike"
  partner_namespace_id        = "pike"
  alias_authorization_rule_id = "pike"
}
