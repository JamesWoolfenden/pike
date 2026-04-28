resource "azurerm_function_app_flex_consumption" "pike_gen" {
  name                = "example-linux-function-app"
  resource_group_name = "pike"
  location            = "pike"
  service_plan_id     = "pike"

  storage_container_type      = "blobContainer"
  storage_container_endpoint  = "${azurerm_storage_account.example.primary_blob_endpoint}${azurerm_storage_container.example.name}"
  storage_authentication_type = "StorageAccountConnectionString"
  runtime_name                = "node"
  runtime_version             = "20"
  maximum_instance_count      = 50
  instance_memory_in_mb       = 2048

  site_config {}
}
