resource "azurerm_storage_account_local_user" "pike_gen" {
  name                 = "user1"
  storage_account_id   = "pike"
  ssh_key_enabled      = true
  ssh_password_enabled = true
  home_directory       = "example_path"
  ssh_authorized_key {
    description = "key1"
    key         = "pike"
  }
  ssh_authorized_key {
    description = "key2"
    key         = "pike"
  }
  permission_scope {
    permissions {
      read   = true
      create = true
    }
    service       = "blob"
    resource_name = "pike"
  }
}
