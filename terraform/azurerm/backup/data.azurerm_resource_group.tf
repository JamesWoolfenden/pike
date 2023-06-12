data "azurerm_resource_group" "pike" {
  name = "pike"
}

output "rg" {
  value = data.azurerm_resource_group.pike
}
