resource "azurerm_storage_blob" "pike_gen" {
  name                   = "my-awesome-content.zip"
  storage_account_name   = "pike"
  storage_container_name = "pike"
  type                   = "Block"
  source                 = "some-local-file.zip"
}
