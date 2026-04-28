resource "azurerm_eventhub_namespace_disaster_recovery_config" "pike_gen" {
  name                 = "replicate-eventhub"
  resource_group_name  = "pike"
  namespace_name       = "pike"
  partner_namespace_id = "pike"
}
