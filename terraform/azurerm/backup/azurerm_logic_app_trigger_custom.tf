resource "azurerm_logic_app_trigger_custom" "pike_gen" {
  name         = "example-trigger"
  logic_app_id = "pike"

  body = <<BODY
{
  "recurrence": {
    "frequency": "Day",
    "interval": 1
  },
  "type": "Recurrence"
}
BODY

}
