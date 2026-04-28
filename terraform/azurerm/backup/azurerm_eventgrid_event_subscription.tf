resource "azurerm_eventgrid_event_subscription" "pike_gen" {
  name  = "example-aees"
  scope = "pike"

  storage_queue_endpoint {
    storage_account_id = "pike"
    queue_name         = "pike"
  }
}
