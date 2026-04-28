resource "azurerm_app_service_source_control_slot" "pike_gen" {
  slot_id  = "pike"
  repo_url = "https://github.com/Azure-Samples/python-docs-hello-world"
  branch   = "master"
}
