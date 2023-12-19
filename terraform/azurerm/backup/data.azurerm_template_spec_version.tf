data "azurerm_template_spec_version" "pike" {
  version             = "1"
  resource_group_name = "pike"
  name                = "pike"
}
