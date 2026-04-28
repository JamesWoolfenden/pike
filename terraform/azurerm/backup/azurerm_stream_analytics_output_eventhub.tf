resource "azurerm_stream_analytics_output_eventhub" "pike_gen" {
  name                      = "output-to-eventhub"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  eventhub_name             = "pike"
  servicebus_namespace      = "pike"
  shared_access_policy_key  = "pike"
  shared_access_policy_name = "RootManageSharedAccessKey"

  serialization {
    type = "Avro"
  }
}
