resource "azurerm_storage_mover_source_endpoint" "pike_gen" {
  name             = "example-se"
  storage_mover_id = "pike"
  export           = "/"
  host             = "192.168.0.1"
  nfs_version      = "NFSv3"
}
