resource "azurerm_network_function_collector_policy" "pike_gen" {
  name                 = "example-nfcp"
  traffic_collector_id = "pike"
  location             = "pike"

  ipfx_emission {
    destination_types = ["AzureMonitor"]
  }

  ipfx_ingestion {
    source_resource_ids = ["pike"]
  }

  tags = {
    key = "value"
  }
}
