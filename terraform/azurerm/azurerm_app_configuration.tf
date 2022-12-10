resource "azurerm_app_configuration" "pike" {
  name                = "appConf1"
  resource_group_name = "pike"
  location            = "uksouth"
}