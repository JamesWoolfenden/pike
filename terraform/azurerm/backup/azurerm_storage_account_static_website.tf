resource "azurerm_storage_account_static_website" "pike_gen" {
  storage_account_id = "pike"
  error_404_document = "custom_not_found.html"
  index_document     = "custom_index.html"
}
