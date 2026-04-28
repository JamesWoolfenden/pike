resource "azurerm_app_service_source_control" "pike_gen" {
  app_id   = "pike"
  repo_url = "https://github.com/Azure-Samples/python-docs-hello-world"
  branch   = "master"
}
