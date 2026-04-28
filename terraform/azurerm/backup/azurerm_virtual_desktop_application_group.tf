resource "azurerm_virtual_desktop_application_group" "pike_gen" {
  name                = "acctag"
  location            = "pike"
  resource_group_name = "pike"

  type          = "RemoteApp"
  host_pool_id  = "pike"
  friendly_name = "TestAppGroup"
  description   = "Acceptance Test: An application group"
}
