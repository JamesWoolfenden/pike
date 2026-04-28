resource "azurerm_storage_mover_target_endpoint" "pike_gen" {
  name                   = "example-se"
  storage_mover_id       = "pike"
  storage_account_id     = "pike"
  storage_container_name = "pike"
  description            = "Example Storage Container Endpoint Description"
}
