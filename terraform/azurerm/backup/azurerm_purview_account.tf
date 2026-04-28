resource "azurerm_purview_account" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"

  identity {
    type = "SystemAssigned"
  }
}
