resource "azurerm_cosmosdb_table" "pike2" {
  name                = "pike2"
  resource_group_name = "pike"
  account_name        = "pike-table"
  autoscale_settings {
    max_throughput = 1000
  }
}


resource "azurerm_cosmosdb_account" "pike-table" {
  name                = "pike-table"
  location            = "uksouth"
  resource_group_name = "pike"
  offer_type          = "Standard"
  enable_free_tier    = true
  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 86400
    max_staleness_prefix    = 1000000
  }

  geo_location {
    location          = "uksouth"
    failover_priority = "0"
    zone_redundant    = false
  }
  capabilities {
    name = "EnableTable"
  }
  tags = {
    "defaultExperience"       = "Azure Table"
    "hidden-cosmos-mmspecial" = ""
  }
}
