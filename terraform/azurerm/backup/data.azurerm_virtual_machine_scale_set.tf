data "azurerm_virtual_machine_scale_set" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
