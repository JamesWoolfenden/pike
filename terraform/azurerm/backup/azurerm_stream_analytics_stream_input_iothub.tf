resource "azurerm_stream_analytics_stream_input_iothub" "pike_gen" {
  name                         = "example-iothub-input"
  stream_analytics_job_name    = "pike"
  resource_group_name          = "pike"
  endpoint                     = "messages/events"
  eventhub_consumer_group_name = "$Default"
  iothub_namespace             = "pike"
  shared_access_policy_key     = "pike"
  shared_access_policy_name    = "iothubowner"

  serialization {
    type     = "Json"
    encoding = "UTF8"
  }
}
