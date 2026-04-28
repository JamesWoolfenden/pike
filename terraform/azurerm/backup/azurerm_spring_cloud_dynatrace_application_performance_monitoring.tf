resource "azurerm_spring_cloud_dynatrace_application_performance_monitoring" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"
  globally_enabled        = true
  api_url                 = "https://example-api-url.com"
  environment_id          = "example-environment-id"
  tenant                  = "example-tenant"
  connection_point        = "https://example.live.dynatrace.com:443"
}
