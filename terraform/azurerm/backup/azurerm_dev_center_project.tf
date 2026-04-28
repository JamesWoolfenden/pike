resource "azurerm_dev_center_project" "pike_gen" {
  dev_center_id       = "pike"
  location            = "pike"
  name                = "example"
  resource_group_name = "pike"
}
