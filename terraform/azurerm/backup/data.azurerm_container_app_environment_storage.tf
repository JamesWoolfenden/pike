data "azurerm_container_app_environment_storage" "pike_gen" {
  name                         = "existing-storage"
  container_app_environment_id = "pike"
}
