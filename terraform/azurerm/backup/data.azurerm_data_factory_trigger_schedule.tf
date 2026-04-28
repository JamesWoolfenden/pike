data "azurerm_data_factory_trigger_schedule" "pike_gen" {
  name            = "example_trigger"
  data_factory_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataFactory/factories/datafactory1"
}
