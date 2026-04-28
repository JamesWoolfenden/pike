resource "azurerm_eventhub" "pike_gen" {
  name              = "acceptanceTestEventHub"
  namespace_id      = "pike"
  partition_count   = 2
  message_retention = 1
}
