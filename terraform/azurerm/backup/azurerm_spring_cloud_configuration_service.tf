resource "azurerm_spring_cloud_configuration_service" "pike_gen" {
  name                    = "default"
  spring_cloud_service_id = "pike"
  repository {
    name                     = "fake"
    label                    = "master"
    patterns                 = ["app/dev"]
    uri                      = "https://github.com/Azure-Samples/piggymetrics"
    search_paths             = ["dir1", "dir2"]
    strict_host_key_checking = false
    username                 = "adminuser"
  }
}
