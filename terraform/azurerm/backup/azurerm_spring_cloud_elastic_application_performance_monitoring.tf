resource "azurerm_spring_cloud_elastic_application_performance_monitoring" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"
  globally_enabled        = true
  application_packages    = ["org.example", "org.another.example"]
  service_name            = "example-service-name"
  server_url              = "http://127.0.0.1:8200"
}
