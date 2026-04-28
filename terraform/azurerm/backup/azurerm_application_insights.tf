resource "azurerm_application_insights" "pike_gen" {
  name                = "tf-test-appinsights"
  location            = "pike"
  resource_group_name = "pike"
  application_type    = "web"
}
