resource "azurerm_spring_cloud_storage" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"
  storage_account_name    = "pike"
  storage_account_key     = "pike"
}
