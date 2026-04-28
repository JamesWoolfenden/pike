resource "azurerm_marketplace_role_assignment" "pike_gen" {
  role_definition_name = "Marketplace Admin"
  principal_id         = "pike"

  lifecycle {
    ignore_changes = [
      name,
      role_definition_id,
    ]
  }
}
