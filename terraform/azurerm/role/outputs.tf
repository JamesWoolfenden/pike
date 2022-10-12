output "password" {
  sensitive = true
  value     = azuread_service_principal_password.pike
}

output "sp" {
  value = azuread_service_principal.pike
}

output "role_definition" {
  value = azurerm_role_definition.example
}
