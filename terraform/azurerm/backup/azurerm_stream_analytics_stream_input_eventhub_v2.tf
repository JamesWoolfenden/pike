resource "azurerm_stream_analytics_stream_input_eventhub_v2" "pike_gen" {
  name                         = "eventhub-stream-input"
  stream_analytics_job_id      = "pike"
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
