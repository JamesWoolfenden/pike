resource "azurerm_virtual_machine_packet_capture" "pike_gen" {
  name               = "example-pc"
  network_watcher_id = "pike"
  virtual_machine_id = "pike"

  storage_location {
    storage_account_id = "pike"
  }

  depends_on = ["pike"]
}
