resource "azurerm_logic_app_integration_account_map" "pike_gen" {
  name                     = "example-iamap"
  resource_group_name      = "pike"
  integration_account_name = "pike"
  map_type                 = "Xslt"
  content                  = file("testdata/integration_account_map_content.xsd")
}
