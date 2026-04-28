resource "azurerm_eventhub_namespace_customer_managed_key" "pike_gen" {
  eventhub_namespace_id = "pike"
  key_vault_key_ids     = ["pike"]
}
