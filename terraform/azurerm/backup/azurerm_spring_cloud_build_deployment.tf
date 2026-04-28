resource "azurerm_spring_cloud_build_deployment" "pike_gen" {
  name                = "example"
  spring_cloud_app_id = "pike"
  build_result_id     = "<default>"
  instance_count      = 2
  environment_variables = {
    "Foo" : "Bar"
    "Env" : "Staging"
  }
  quota {
    cpu    = "2"
    memory = "4Gi"
  }
}
