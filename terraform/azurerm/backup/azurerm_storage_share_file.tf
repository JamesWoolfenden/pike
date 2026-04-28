resource "azurerm_storage_share_file" "pike_gen" {
  name              = "my-awesome-content.zip"
  storage_share_url = "pike"
  source            = "some-local-file.zip"
}
