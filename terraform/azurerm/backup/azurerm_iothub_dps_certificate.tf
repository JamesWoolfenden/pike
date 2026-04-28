resource "azurerm_iothub_dps_certificate" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  iot_dps_name        = "pike"

  certificate_content = filebase64("example.cer")
}
