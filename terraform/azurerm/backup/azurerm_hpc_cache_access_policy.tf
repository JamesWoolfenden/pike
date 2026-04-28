resource "azurerm_hpc_cache_access_policy" "pike_gen" {
  name         = "example"
  hpc_cache_id = "pike"

  access_rule {
    scope  = "default"
    access = "rw"
  }
}
