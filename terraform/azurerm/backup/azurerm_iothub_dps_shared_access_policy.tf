resource "azurerm_iothub_dps_shared_access_policy" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  iothub_dps_name     = "pike"

  enrollment_write = true
  enrollment_read  = true
}
