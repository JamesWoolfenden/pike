resource "azurerm_hpc_cache_nfs_target" "pike_gen" {
  name                = "examplehpcnfstarget"
  resource_group_name = "pike"
  cache_name          = "pike"
  target_host_name    = "pike"
  usage_model         = "READ_HEAVY_INFREQ"
  namespace_junction {
    namespace_path = "/nfs/a1"
    nfs_export     = "/export/a"
    target_path    = "1"
  }
  namespace_junction {
    namespace_path = "/nfs/b"
    nfs_export     = "/export/b"
  }
}
