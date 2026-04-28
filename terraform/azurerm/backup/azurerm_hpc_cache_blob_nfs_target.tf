resource "azurerm_hpc_cache_blob_nfs_target" "pike_gen" {
  name                 = "example-hpc-target"
  resource_group_name  = "pike"
  cache_name           = "pike"
  storage_container_id = "pike"
  namespace_path       = "/p1"
  usage_model          = "READ_HEAVY_INFREQ"
}
