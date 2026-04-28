data "azurerm_role_management_policy" "pike_gen" {
  scope              = "pike"
  role_definition_id = "pike"
}
