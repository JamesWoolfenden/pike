resource "azurerm_servicebus_queue" "pike_gen" {
  name         = "tfex_servicebus_queue"
  namespace_id = "pike"

  partitioning_enabled = true
}
