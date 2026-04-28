resource "azurerm_dev_center_catalog" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  dev_center_id       = "pike"
  catalog_github {
    branch            = "foo"
    path              = ""
    uri               = "example URI"
    key_vault_key_url = "secret"
  }
}
