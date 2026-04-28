resource "azurerm_automation_module" "pike_gen" {
  name                    = "xActiveDirectory"
  resource_group_name     = "pike"
  automation_account_name = "pike"

  module_link {
    uri = "https://devopsgallerystorage.blob.core.windows.net/packages/xactivedirectory.2.19.0.nupkg"
  }
}
