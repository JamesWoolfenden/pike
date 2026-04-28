resource "azurerm_digital_twins_endpoint_eventgrid" "pike_gen" {
  name                     = "example-EG"
  digital_twins_id         = "pike"
  eventgrid_topic_endpoint = "pike"
}
