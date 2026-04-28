resource "azurerm_data_factory_linked_service_key_vault" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  key_vault_id    = "pike"
}
