resource "azurerm_arc_private_link_scope" "pike_gen" {
  name                = "plsexample"
  resource_group_name = "pike"
  location            = "pike"
}
