resource "azurerm_traffic_manager_profile" "pike_gen" {
  name                   = "pike"
  resource_group_name    = "pike"
  traffic_routing_method = "Weighted"

  dns_config {
    relative_name = "pike"
    ttl           = 100
  }

  monitor_config {
    protocol                     = "HTTP"
    port                         = 80
    path                         = "/"
    interval_in_seconds          = 30
    timeout_in_seconds           = 9
    tolerated_number_of_failures = 3
  }

  tags = {
    environment = "Production"
  }
}
