resource "azurerm_stream_analytics_job_storage_account" "pike_gen" {
  stream_analytics_job_id = "pike"
  storage_account_name    = "pike"
  authentication_mode     = "Msi"
}
