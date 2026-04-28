resource "azurerm_iothub_device_update_instance" "pike_gen" {
  name                     = "example"
  device_update_account_id = "pike"
  iothub_id                = "pike"
  diagnostic_enabled       = true

  diagnostic_storage_account {
    id = "pike"
  }

  tags = {
    key = "value"
  }
}
