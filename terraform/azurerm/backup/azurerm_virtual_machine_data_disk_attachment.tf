resource "azurerm_virtual_machine_data_disk_attachment" "pike_gen" {
  managed_disk_id    = "pike"
  virtual_machine_id = "pike"
  lun                = "10"
  caching            = "ReadWrite"
}
