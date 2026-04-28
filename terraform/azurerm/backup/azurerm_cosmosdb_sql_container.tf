resource "azurerm_cosmosdb_sql_container" "pike_gen" {
  name                  = "example-container"
  resource_group_name   = "pike"
  account_name          = "pike"
  database_name         = "pike"
  partition_key_paths   = ["/definition/id"]
  partition_key_version = 1
  throughput            = 400

  indexing_policy {
    indexing_mode = "consistent"

    included_path {
      path = "/*"
    }

    included_path {
      path = "/included/?"
    }

    excluded_path {
      path = "/excluded/?"
    }
  }

  unique_key {
    paths = ["/definition/idlong", "/definition/idshort"]
  }
}
