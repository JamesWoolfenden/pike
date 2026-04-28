resource "azurerm_iot_security_device_group" "pike_gen" {
  name      = "example-device-security-group"
  iothub_id = "pike"

  allow_rule {
    connection_to_ips_not_allowed = ["10.0.0.0/24"]
  }

  range_rule {
    type     = "ActiveConnectionsNotInAllowedRange"
    min      = 0
    max      = 30
    duration = "PT5M"
  }

  depends_on = ["pike"]
}
