resource "azurerm_application_insights_smart_detection_rule" "pike_gen" {
  name                    = "Slow server response time"
  application_insights_id = "pike"
  enabled                 = false
}
