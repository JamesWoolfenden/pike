resource "azurerm_advanced_threat_protection" "pike_gen" {
  target_resource_id = "pike"
  enabled            = true
}
