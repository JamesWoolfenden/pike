resource "azurerm_key_vault_managed_storage_account" "pike_gen" {
  name                         = "examplemanagedstorage"
  key_vault_id                 = "pike"
  storage_account_id           = "pike"
  storage_account_key          = "key1"
  regenerate_key_automatically = false
  regeneration_period          = "P1D"
}
