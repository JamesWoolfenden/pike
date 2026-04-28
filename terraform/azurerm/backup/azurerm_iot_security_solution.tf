resource "azurerm_iot_security_solution" "pike_gen" {
  name                = "example-Iot-Security-Solution"
  resource_group_name = "pike"
  location            = "pike"
  display_name        = "Iot Security Solution"
  iothub_ids          = ["pike"]
}
