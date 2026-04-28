resource "azurerm_stream_analytics_output_synapse" "pike_gen" {
  name                      = "example-output-synapse"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"

  server   = "pike" ["sqlOnDemand"]
  user     = "pike"
  database = "master"
  table    = "ExampleTable"
}
