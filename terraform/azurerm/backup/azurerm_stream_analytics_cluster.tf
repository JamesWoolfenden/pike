resource "azurerm_stream_analytics_cluster" "pike_gen" {
  name                = "examplestreamanalyticscluster"
  resource_group_name = "pike"
  location            = "pike"
  streaming_capacity  = 36
}
