resource "azurerm_stream_analytics_output_servicebus_queue" "pike_gen" {
  name                      = "blob-storage-output"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  queue_name                = "pike"
  servicebus_namespace      = "pike"
  shared_access_policy_key  = "pike"
  shared_access_policy_name = "RootManageSharedAccessKey"

  serialization {
    type   = "Csv"
    format = "Array"
  }
}
