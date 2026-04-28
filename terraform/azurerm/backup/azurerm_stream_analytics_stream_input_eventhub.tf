resource "azurerm_stream_analytics_stream_input_eventhub" "pike_gen" {
  name                         = "eventhub-stream-input"
  stream_analytics_job_name    = "pike"
  resource_group_name          = "pike"
  eventhub_consumer_group_name = "pike"
  eventhub_name                = "pike"
  servicebus_namespace         = "pike"
  shared_access_policy_key     = "pike"
  shared_access_policy_name    = "RootManageSharedAccessKey"

  serialization {
    type     = "Json"
    encoding = "UTF8"
  }
}
