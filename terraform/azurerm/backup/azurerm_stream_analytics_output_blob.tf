resource "azurerm_stream_analytics_output_blob" "pike_gen" {
  name                      = "output-to-blob-storage"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  storage_account_name      = "pike"
  storage_account_key       = "pike"
  storage_container_name    = "pike"
  path_pattern              = "some-pattern"
  date_format               = "yyyy-MM-dd"
  time_format               = "HH"

  serialization {
    type            = "Csv"
    encoding        = "UTF8"
    field_delimiter = ","
  }
}
