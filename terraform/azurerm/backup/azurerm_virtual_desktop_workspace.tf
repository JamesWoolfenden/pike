resource "azurerm_virtual_desktop_workspace" "pike_gen" {
  name                = "workspace"
  location            = "pike"
  resource_group_name = "pike"

  friendly_name = "FriendlyName"
  description   = "A description of my workspace"
}
