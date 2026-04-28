resource "azurerm_service_fabric_cluster" "pike_gen" {
  name                 = "example-servicefabric"
  resource_group_name  = "pike"
  location             = "pike"
  reliability_level    = "Bronze"
  upgrade_mode         = "Manual"
  cluster_code_version = "7.1.456.959"
  vm_image             = "Windows"
  management_endpoint  = "https://example:80"

  node_type {
    name                 = "first"
    instance_count       = 3
    is_primary           = true
    client_endpoint_port = 2020
    http_endpoint_port   = 80
  }
}
