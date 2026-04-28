resource "azurerm_kusto_database" "pike_gen" {
  name                = "my-kusto-database"
  resource_group_name = "pike"
  location            = "pike"
  cluster_name        = "pike"

  hot_cache_period   = "P7D"
  soft_delete_period = "P31D"

  # prevent the possibility of accidental data loss
  lifecycle {
    prevent_destroy = true
  }
}
