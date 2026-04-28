resource "azurerm_hpc_cache_blob_target" "pike_gen" {
  name                 = "examplehpccblobtarget"
  resource_group_name  = "pike"
  cache_name           = "pike"
  storage_container_id = "pike"
  namespace_path       = "/blob_storage"
}
