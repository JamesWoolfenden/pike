resource "azurerm_container_app" "pike_gen" {
  name                         = "example-app"
  container_app_environment_id = "pike"
  resource_group_name          = "pike"
  revision_mode                = "Single"

  template {
    container {
      name   = "examplecontainerapp"
      image  = "mcr.microsoft.com/k8se/quickstart:latest"
      cpu    = 0.25
      memory = "0.5Gi"
    }
  }
}
