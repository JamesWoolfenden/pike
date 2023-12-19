data "azurerm_api_management_api" "pike" {
  name                = "pike"
  revision            = "1"
  resource_group_name = "pike"
  api_management_name = "pike"
}
