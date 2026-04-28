resource "azurerm_container_registry_token_password" "pike_gen" {
  container_registry_token_id = "pike"

  password1 {
    expiry = "2023-03-22T17:57:36+08:00"
  }
}
