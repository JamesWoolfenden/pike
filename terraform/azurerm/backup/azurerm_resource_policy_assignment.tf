resource "azurerm_resource_policy_assignment" "pike_gen" {
  name                 = "example-policy-assignment"
  resource_id          = "pike"
  policy_definition_id = "pike"
}
