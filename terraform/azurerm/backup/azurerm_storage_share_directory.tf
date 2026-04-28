resource "azurerm_storage_share_directory" "pike_gen" {
  name              = "example"
  storage_share_url = "pike"
}
