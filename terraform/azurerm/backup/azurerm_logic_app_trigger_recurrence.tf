resource "azurerm_logic_app_trigger_recurrence" "pike_gen" {
  name         = "run-every-day"
  logic_app_id = "pike"
  frequency    = "Day"
  interval     = 1
}
