data "azurerm_managed_disk" "pike_gen" {
  name                = "example-datadisk"
  resource_group_name = "example-resources"
}
