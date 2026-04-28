resource "azurerm_cosmosdb_gremlin_graph" "pike_gen" {
  name                = "tfex-cosmos-gremlin-graph"
  resource_group_name = "pike"
  account_name        = "pike"
  database_name       = "pike"
  partition_key_path  = "/Example"
  throughput          = 400

  index_policy {
    automatic      = true
    indexing_mode  = "consistent"
    included_paths = ["/*"]
    excluded_paths = ["/\"_etag\"/?"]
  }

  conflict_resolution_policy {
    mode                     = "LastWriterWins"
    conflict_resolution_path = "/_ts"
  }

  unique_key {
    paths = ["/definition/id1", "/definition/id2"]
  }
}
