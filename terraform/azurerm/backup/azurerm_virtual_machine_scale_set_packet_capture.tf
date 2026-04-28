resource "azurerm_virtual_machine_scale_set_packet_capture" "pike_gen" {
  name                         = "example-pc"
  network_watcher_id           = "pike"
  virtual_machine_scale_set_id = "pike"

  storage_location {
    file_path = "/var/captures/packet.cap"
  }

  machine_scope {
    include_instance_ids = ["0"]
    exclude_instance_ids = ["1"]
  }

  depends_on = ["pike"]
}
