resource "azurerm_iothub_fallback_route" "pike_gen" {
  resource_group_name = "pike"
  iothub_name         = "pike"

  condition      = "true"
  endpoint_names = ["pike"]
  enabled        = true
}
