resource "azurerm_servicebus_topic" "pike_gen" {
  name         = "tfex_servicebus_topic"
  namespace_id = "pike"

  partitioning_enabled = true
}
