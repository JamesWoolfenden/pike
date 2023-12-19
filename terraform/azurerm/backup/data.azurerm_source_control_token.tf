data "azurerm_source_control_token" "pike" {
  type = "GitHub"
}

output "token" {
  value     = data.azurerm_source_control_token.pike
  sensitive = true
}
