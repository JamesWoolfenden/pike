resource "azurerm_stream_analytics_stream_input_blob" "pike_gen" {
  name                      = "blob-stream-input"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  storage_account_name      = "pike"
  storage_account_key       = "pike"
  storage_container_name    = "pike"
  path_pattern              = "some-random-pattern"
  date_format               = "yyyy/MM/dd"
  time_format               = "HH"

  serialization {
    type     = "Json"
    encoding = "UTF8"
  }
}
