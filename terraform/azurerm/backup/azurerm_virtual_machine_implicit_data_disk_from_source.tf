resource "azurerm_virtual_machine_implicit_data_disk_from_source" "pike_gen" {
  name               = "${local.vm_name}-implicitdisk1"
  virtual_machine_id = "pike"
  lun                = "0"
  caching            = "None"
  create_option      = "Copy"
  disk_size_gb       = 20
  source_resource_id = "pike"
}
