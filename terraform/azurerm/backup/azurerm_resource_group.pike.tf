resource "azurerm_resource_group" "pike" {
  name     = "pike"
  location = "West Europe"
  tags = {
    pike   = "permission"
    delete = "me"
  }
}
