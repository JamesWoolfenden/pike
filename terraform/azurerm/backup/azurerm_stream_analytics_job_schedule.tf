resource "azurerm_stream_analytics_job_schedule" "pike_gen" {
  stream_analytics_job_id = "pike"
  start_mode              = "CustomTime"
  start_time              = "2022-09-21T00:00:00Z"

  depends_on = [
    azurerm_stream_analytics_job.example,
    azurerm_stream_analytics_stream_input_blob.example,
    azurerm_stream_analytics_output_blob.example,
  ]
}
