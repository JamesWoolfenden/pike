resource "azurerm_eventgrid_namespace_topic" "pike_gen" {
  name                    = "topic-namespace-example"
  eventgrid_namespace_id  = "pike"
  event_retention_in_days = 1
}
