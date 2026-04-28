resource "azurerm_storage_mover_project" "pike_gen" {
  name             = "example-sp"
  storage_mover_id = "pike"
  description      = "Example Project Description"
}
