resource "azurerm_management_group_policy_assignment" "pike_gen" {
  name                 = "example-policy"
  policy_definition_id = "pike"
  management_group_id  = "pike"
}
