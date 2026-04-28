resource "azurerm_storage_encryption_scope" "pike_gen" {
  name               = "microsoftmanaged"
  storage_account_id = "pike"
  source             = "Microsoft.Storage"
}
