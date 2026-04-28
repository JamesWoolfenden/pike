resource "azurerm_automation_powershell72_module" "pike_gen" {
  name                  = "xActiveDirectory"
  automation_account_id = "pike"

  module_link {
    uri = "https://devopsgallerystorage.blob.core.windows.net/packages/xactivedirectory.2.19.0.nupkg"
  }
}
