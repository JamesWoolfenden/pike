resource "azurerm_stream_analytics_output_servicebus_topic" "pike_gen" {
  name                      = "service-bus-topic-output"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  topic_name                = "pike"
  servicebus_namespace      = "pike"
  shared_access_policy_key  = "pike"
  shared_access_policy_name = "RootManageSharedAccessKey"
  property_columns          = ["col1", "col2"]

  serialization {
    type   = "Csv"
    format = "Array"
  }
}
