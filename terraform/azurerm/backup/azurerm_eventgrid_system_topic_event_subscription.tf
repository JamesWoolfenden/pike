resource "azurerm_eventgrid_system_topic_event_subscription" "pike_gen" {
  name                = "example-event-subscription"
  system_topic        = "pike"
  resource_group_name = "pike"

  storage_queue_endpoint {
    storage_account_id = "pike"
    queue_name         = "pike"
  }
}
