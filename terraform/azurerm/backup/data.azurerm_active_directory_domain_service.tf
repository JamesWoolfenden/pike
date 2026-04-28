data "azurerm_active_directory_domain_service" "pike_gen" {
  name                = "example-aadds"
  resource_group_name = "example-aadds-rg"
}
