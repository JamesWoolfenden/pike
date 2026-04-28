resource "azurerm_healthcare_workspace" "pike_gen" {
  name                = "tfexworkspace"
  resource_group_name = "tfex-resource_group"
  location            = "east us"
}
