resource "azurerm_virtual_machine_restore_point" "pike_gen" {
  name                                        = "example-restore-point"
  virtual_machine_restore_point_collection_id = "pike"
}
