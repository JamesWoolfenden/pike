resource "azurerm_powerbi_embedded" "pike_gen" {
  name                = "examplepowerbi"
  location            = "pike"
  resource_group_name = "pike"
  sku_name            = "A1"
  administrators      = ["azsdktest@microsoft.com"]
}
