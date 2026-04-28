resource "azurerm_virtual_hub_routing_intent" "pike_gen" {
  name           = "example-routingintent"
  virtual_hub_id = "pike"

  routing_policy {
    name         = "InternetTrafficPolicy"
    destinations = ["Internet"]
    next_hop     = "pike"
  }
}
