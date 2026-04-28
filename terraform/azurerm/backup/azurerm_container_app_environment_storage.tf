resource "azurerm_container_app_environment_storage" "pike_gen" {
  name                         = "mycontainerappstorage"
  container_app_environment_id = "pike"
  account_name                 = "pike"
  share_name                   = "pike"
  access_mode                  = "ReadOnly"
}
