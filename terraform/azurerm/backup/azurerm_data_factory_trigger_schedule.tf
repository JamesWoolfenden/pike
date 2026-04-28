resource "azurerm_data_factory_trigger_schedule" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  pipeline_name   = "pike"

  interval  = 5
  frequency = "Day"
}
