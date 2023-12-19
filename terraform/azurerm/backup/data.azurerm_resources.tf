data "azurerm_resources" "pike" {
  name                = "pike"
  resource_group_name = "pike"
}

output "resources" {
  value = data.azurerm_resources.pike
}
