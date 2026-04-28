data "azurerm_servicebus_namespace_disaster_recovery_config" "pike_gen" {
  name         = "existing"
  namespace_id = "example-namespace-id"
}
