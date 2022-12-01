resource "azurerm_api_management" "pike" {
  name                = "example-pike2"
  location            = "uksouth"
  resource_group_name = "pike"
  publisher_name      = "My Company"
  publisher_email     = "company@terraform.io"

  sku_name = "Developer_1"

}
