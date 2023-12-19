data "azurerm_bastion_host" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
