resource "azurerm_maintenance_assignment_dedicated_host" "pike_gen" {
  location                     = "pike"
  maintenance_configuration_id = "pike"
  dedicated_host_id            = "pike"
}
