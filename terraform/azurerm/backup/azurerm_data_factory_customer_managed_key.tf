resource "azurerm_data_factory_customer_managed_key" "pike_gen" {
  data_factory_id         = "pike"
  customer_managed_key_id = "pike"
}
