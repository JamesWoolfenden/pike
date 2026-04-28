resource "azurerm_servicebus_namespace_customer_managed_key" "pike_gen" {
  namespace_id     = "pike"
  key_vault_key_id = "pike"
}
