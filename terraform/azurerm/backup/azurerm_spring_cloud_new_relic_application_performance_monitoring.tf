resource "azurerm_spring_cloud_new_relic_application_performance_monitoring" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"
  app_name                = "example-app-name"
  license_key             = "example-license-key"
  app_server_port         = 8080
  labels = {
    tagName1 = "tagValue1"
    tagName2 = "tagValue2"
  }
  globally_enabled = true
}
