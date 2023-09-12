data "azurerm_app_service_plan" "pike" {
  resource_group_name = "pike"
  name                = azurerm_app_service_plan.pike.name
}

output "plan" {
  value = data.azurerm_app_service_plan.pike
}
