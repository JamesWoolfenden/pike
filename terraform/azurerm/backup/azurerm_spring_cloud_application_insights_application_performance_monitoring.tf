resource "azurerm_spring_cloud_application_insights_application_performance_monitoring" "pike_gen" {
  name                         = "example"
  spring_cloud_service_id      = "pike"
  globally_enabled             = true
  role_name                    = "test-role"
  role_instance                = "test-instance"
  sampling_percentage          = 50
  sampling_requests_per_second = 10
}
