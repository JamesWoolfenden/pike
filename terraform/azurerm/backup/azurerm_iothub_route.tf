resource "azurerm_iothub_route" "pike_gen" {
  resource_group_name = "pike"
  iothub_name         = "pike"
  name                = "example"

  source         = "DeviceMessages"
  condition      = "true"
  endpoint_names = ["pike"]
  enabled        = true
}
