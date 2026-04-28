resource "azurerm_stream_analytics_output_powerbi" "pike_gen" {
  name                    = "output-to-powerbi"
  stream_analytics_job_id = "pike"
  dataset                 = "example-dataset"
  table                   = "example-table"
  group_id                = "00000000-0000-0000-0000-000000000000"
  group_name              = "some-group-name"
}
