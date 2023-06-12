data "azurerm_client_config" "current" {}

output "config" {
  value = data.azurerm_client_config.current
}
