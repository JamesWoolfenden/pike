data "azurerm_storage_blob" "pike" {
  storage_container_name = "pike"
  storage_account_name   = "pike"
  name                   = "pike"
}
