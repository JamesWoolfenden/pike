resource "azurerm_dashboard_grafana_managed_private_endpoint" "pike_gen" {
  grafana_id                   = "pike"
  name                         = "example-mpe"
  location                     = "pike"
  private_link_resource_id     = "pike"
  group_ids                    = ["prometheusMetrics"]
  private_link_resource_region = "pike"
}
